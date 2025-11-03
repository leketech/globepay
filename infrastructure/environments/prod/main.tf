# infrastructure/environments/prod/main.tf

terraform {
  required_version = ">= 1.5.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.23"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.11"
    }
  }

  backend "s3" {
    bucket         = "globepay-terraform-state-prod"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "globepay-terraform-locks"
  }
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project     = "Globepay"
      Environment = var.environment
      ManagedBy   = "Terraform"
      CostCenter  = "Engineering"
    }
  }
}

locals {
  cluster_name = "${var.project_name}-${var.environment}-eks"
  common_tags = {
    Project     = var.project_name
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}

# VPC and Networking
module "networking" {
  source = "../../modules/networking"

  project_name        = var.project_name
  environment         = var.environment
  vpc_cidr            = var.vpc_cidr
  availability_zones  = var.availability_zones
  public_subnet_cidrs = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  database_subnet_cidrs = var.database_subnet_cidrs
  enable_nat_gateway  = true
  single_nat_gateway  = false
  enable_vpn_gateway  = false
  enable_dns_hostnames = true
  enable_dns_support  = true
  
  tags = local.common_tags
}

# KMS for encryption
module "kms" {
  source = "../../modules/kms"

  project_name = var.project_name
  environment  = var.environment
  
  tags = local.common_tags
}

# EKS Cluster
module "eks" {
  source = "../../modules/eks"

  project_name                 = var.project_name
  environment                  = var.environment
  cluster_name                 = local.cluster_name
  cluster_version              = var.eks_cluster_version
  vpc_id                       = module.networking.vpc_id
  private_subnet_ids           = module.networking.private_subnet_ids
  cluster_endpoint_public_access = var.cluster_endpoint_public_access
  cluster_endpoint_private_access = true
  cluster_endpoint_public_access_cidrs = ["102.22.168.11/32"]
  
  # Node groups
  node_groups = {
    general = {
      desired_size   = 2
      min_size       = 1
      max_size       = 3
      instance_types = ["t3.large"]
      capacity_type  = "ON_DEMAND"
      disk_size      = 50
      labels = {
        role = "general"
      }
      taints = []
      tags   = {}
    }
    
    api = {
      desired_size   = 2
      min_size       = 1
      max_size       = 3
      instance_types = ["t3.xlarge"]
      capacity_type  = "ON_DEMAND"
      disk_size      = 50
      labels = {
        role = "api"
        workload = "backend"
      }
      taints = []
      tags   = {}
    }
    
    worker = {
      desired_size   = 1
      min_size       = 1
      max_size       = 2
      instance_types = ["t3.medium"]
      capacity_type  = "SPOT"
      disk_size      = 50
      labels = {
        role = "worker"
        workload = "background-jobs"
      }
      taints = []
      tags   = {}
    }
  }

  encryption_key_arn = module.kms.eks_key_arn
  
  tags = local.common_tags
}

# RDS PostgreSQL
module "rds" {
  source = "../../modules/rds"

  project_name              = var.project_name
  environment               = var.environment
  vpc_id                    = module.networking.vpc_id
  db_subnet_group_name      = module.networking.db_subnet_group_name
  allowed_security_group_ids = [module.eks.cluster_security_group_id]
  
  # Database configuration
  engine                    = "postgres"
  engine_version            = "15.12"
  instance_class            = "db.t4g.medium"
  allocated_storage         = 500
  max_allocated_storage     = 2000
  storage_encrypted         = true
  kms_key_id                = module.kms.rds_key_arn
  
  # High availability
  multi_az                  = true
  backup_retention_period   = 30
  backup_window             = "03:00-04:00"
  maintenance_window        = "Mon:04:00-Mon:05:00"
  
  # Performance
  performance_insights_enabled = true
  monitoring_interval       = 60
  enabled_cloudwatch_logs_exports = ["postgresql", "upgrade"]
  
  # Security
  deletion_protection       = true
  skip_final_snapshot       = false
  final_snapshot_identifier = "${var.project_name}-${var.environment}-final-snapshot"
  
  # Use a custom DB instance identifier to avoid conflict
  db_instance_identifier    = "${var.project_name}-${var.environment}-db-2"
  
  database_name = var.database_name
  master_username = var.database_username
  master_password = var.database_password
  
  tags = local.common_tags
}

# ElastiCache Redis
module "elasticache" {
  source = "../../modules/elasticache"

  project_name              = var.project_name
  environment               = var.environment
  vpc_id                    = module.networking.vpc_id
  subnet_ids                = module.networking.private_subnet_ids
  allowed_security_group_ids = [module.eks.cluster_security_group_id]
  
  # Redis configuration
  engine_version            = "7.0"
  node_type                 = "cache.t4g.medium"
  num_cache_nodes           = 2
  parameter_group_family    = "redis7"
  
  # High availability
  automatic_failover_enabled = true
  multi_az_enabled          = true
  
  # Backup
  snapshot_retention_limit  = 7
  snapshot_window           = "03:00-04:00"
  
  # Encryption
  at_rest_encryption_enabled = true
  transit_encryption_enabled = true
  auth_token_enabled        = true
  kms_key_id                = module.kms.elasticache_key_arn
  
  tags = local.common_tags
}

# S3 Buckets
module "s3" {
  source = "../../modules/s3"

  project_name = var.project_name
  environment  = var.environment
  
  buckets = {
    documents = {
      versioning_enabled = true
      lifecycle_rules = [
        {
          id      = "archive-old-documents"
          enabled = true
          transition = {
            days          = 90
            storage_class = "GLACIER"
          }
        }
      ]
    }
    
    backups = {
      versioning_enabled = true
      lifecycle_rules = [
        {
          id      = "delete-old-backups"
          enabled = true
          expiration = {
            days = 365
          }
        }
      ]
    }
    
    logs = {
      versioning_enabled = false
      lifecycle_rules = [
        {
          id      = "delete-old-logs"
          enabled = true
          expiration = {
            days = 90
          }
        }
      ]
    }
    
    assets = {
      versioning_enabled = true
      lifecycle_rules    = []
    }
  }
  
  kms_key_id = module.kms.s3_key_arn
  tags       = local.common_tags
}

# CloudFront for Frontend
module "cloudfront" {
  source = "../../modules/cloudfront"

  project_name    = var.project_name
  environment     = var.environment
  domain_name     = var.domain_name
  s3_bucket_id    = module.s3.bucket_ids["assets"]
  s3_bucket_domain_name = module.s3.bucket_domain_names["assets"]
  s3_bucket_arn   = module.s3.bucket_arns["assets"]
  acm_certificate_arn = var.acm_certificate_arn
  
  tags = local.common_tags
}

# SES for Email
module "ses" {
  source = "../../modules/ses"

  project_name = var.project_name
  environment  = var.environment
  domain_name  = var.domain_name
  route53_zone_id = var.route53_zone_id
  verify_domain   = false
  
  tags = local.common_tags
}

# Monitoring
module "monitoring" {
  source = "../../modules/monitoring"

  project_name        = var.project_name
  environment         = var.environment
  eks_cluster_name    = module.eks.cluster_name
  rds_instance_id     = module.rds.db_instance_id
  elasticache_cluster_id = module.elasticache.cluster_id
  
  # SNS topic for alerts
  alert_email = var.alert_email
  
  tags = local.common_tags
}

# Outputs have been moved to outputs.tf
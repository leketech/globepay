# Terraform configuration for development environment

terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# AWS Provider
provider "aws" {
  region = var.aws_region
  # Add your AWS credentials configuration here
  # You can use AWS credentials file, environment variables, or IAM roles
}

# Include modules
module "networking" {
  source = "../../modules/networking"
  
  environment = "dev"
  vpc_cidr    = "10.0.0.0/16"
  azs         = var.azs
}

module "eks" {
  source = "../../modules/eks"
  
  environment     = "dev"
  vpc_id          = module.networking.vpc_id
  private_subnets = module.networking.private_subnets
  eks_version     = "1.28"
}

module "rds" {
  source = "../../modules/rds"
  
  environment     = "dev"
  vpc_id          = module.networking.vpc_id
  private_subnets = module.networking.private_subnets
  db_name         = "globepay_dev"
  db_username     = "globepay"
  db_password     = var.db_password
}

module "elasticache" {
  source = "../../modules/elasticache"
  
  environment     = "dev"
  vpc_id          = module.networking.vpc_id
  private_subnets = module.networking.private_subnets
}

module "s3" {
  source = "../../modules/s3"
  
  environment = "dev"
}

module "cloudfront" {
  source = "../../modules/cloudfront"
  
  environment    = "dev"
  s3_bucket_name = module.s3.bucket_name
}

module "ses" {
  source = "../../modules/ses"
  
  environment = "dev"
}

module "kms" {
  source = "../../modules/kms"
  
  environment = "dev"
}

module "monitoring" {
  source = "../../modules/monitoring"
  
  environment = "dev"
}
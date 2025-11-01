# Outputs for development environment

output "vpc_id" {
  description = "VPC ID"
  value       = module.networking.vpc_id
}

output "public_subnet_ids" {
  description = "Public subnet IDs"
  value       = module.networking.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Private subnet IDs"
  value       = module.networking.private_subnet_ids
}

output "database_subnet_ids" {
  description = "Database subnet IDs"
  value       = module.networking.database_subnet_ids
}

output "eks_cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = module.eks.cluster_endpoint
  sensitive   = true
}

output "eks_cluster_name" {
  description = "EKS cluster name"
  value       = module.eks.cluster_name
}

output "eks_cluster_certificate_authority_data" {
  description = "EKS cluster certificate authority data"
  value       = module.eks.cluster_certificate_authority_data
  sensitive   = true
}

output "rds_endpoint" {
  description = "RDS endpoint"
  value       = module.rds.db_instance_endpoint
  sensitive   = true
}

output "rds_database_name" {
  description = "RDS database name"
  value       = module.rds.db_instance_name
}

output "redis_endpoint" {
  description = "Redis endpoint"
  value       = module.elasticache.configuration_endpoint
  sensitive   = true
}

output "s3_bucket_names" {
  description = "S3 bucket names"
  value       = module.s3.bucket_names
}

output "s3_bucket_arns" {
  description = "S3 bucket ARNs"
  value       = module.s3.bucket_arns
}

output "cloudfront_distribution_id" {
  description = "CloudFront distribution ID"
  value       = module.cloudfront.distribution_id
}

output "cloudfront_domain_name" {
  description = "CloudFront domain name"
  value       = module.cloudfront.distribution_domain_name
}

output "ses_domain_identity_arn" {
  description = "SES domain identity ARN"
  value       = module.ses.domain_identity_arn
}

output "kms_key_arns" {
  description = "KMS key ARNs"
  value = {
    eks       = module.kms.eks_key_arn
    rds       = module.kms.rds_key_arn
    elasticache = module.kms.elasticache_key_arn
    s3        = module.kms.s3_key_arn
  }
  sensitive = true
}
# Outputs for KMS module

output "eks_key_id" {
  description = "EKS KMS key ID"
  value       = aws_kms_key.eks.key_id
}

output "eks_key_arn" {
  description = "EKS KMS key ARN"
  value       = aws_kms_key.eks.arn
}

output "rds_key_id" {
  description = "RDS KMS key ID"
  value       = aws_kms_key.rds.key_id
}

output "rds_key_arn" {
  description = "RDS KMS key ARN"
  value       = aws_kms_key.rds.arn
}

output "elasticache_key_id" {
  description = "ElastiCache KMS key ID"
  value       = aws_kms_key.elasticache.key_id
}

output "elasticache_key_arn" {
  description = "ElastiCache KMS key ARN"
  value       = aws_kms_key.elasticache.arn
}

output "s3_key_id" {
  description = "S3 KMS key ID"
  value       = aws_kms_key.s3.key_id
}

output "s3_key_arn" {
  description = "S3 KMS key ARN"
  value       = aws_kms_key.s3.arn
}
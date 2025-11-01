# KMS Keys Module

# EKS KMS Key
resource "aws_kms_key" "eks" {
  description             = "KMS key for EKS encryption - ${var.project_name} ${var.environment}"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-eks-key"
    }
  )
}

resource "aws_kms_alias" "eks" {
  name          = "alias/${var.project_name}-${var.environment}-eks-key"
  target_key_id = aws_kms_key.eks.key_id
}

# RDS KMS Key
resource "aws_kms_key" "rds" {
  description             = "KMS key for RDS encryption - ${var.project_name} ${var.environment}"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-rds-key"
    }
  )
}

resource "aws_kms_alias" "rds" {
  name          = "alias/${var.project_name}-${var.environment}-rds-key"
  target_key_id = aws_kms_key.rds.key_id
}

# ElastiCache KMS Key
resource "aws_kms_key" "elasticache" {
  description             = "KMS key for ElastiCache encryption - ${var.project_name} ${var.environment}"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-elasticache-key"
    }
  )
}

resource "aws_kms_alias" "elasticache" {
  name          = "alias/${var.project_name}-${var.environment}-elasticache-key"
  target_key_id = aws_kms_key.elasticache.key_id
}

# S3 KMS Key
resource "aws_kms_key" "s3" {
  description             = "KMS key for S3 encryption - ${var.project_name} ${var.environment}"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-s3-key"
    }
  )
}

resource "aws_kms_alias" "s3" {
  name          = "alias/${var.project_name}-${var.environment}-s3-key"
  target_key_id = aws_kms_key.s3.key_id
}
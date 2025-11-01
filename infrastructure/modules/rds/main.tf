# RDS PostgreSQL Module

resource "aws_db_instance" "main" {
  identifier              = var.db_instance_identifier != null ? var.db_instance_identifier : "${var.project_name}-${var.environment}-db"
  db_name                 = var.database_name
  username                = var.master_username
  password                = var.master_password
  engine                  = var.engine
  engine_version          = var.engine_version
  instance_class          = var.instance_class
  allocated_storage       = var.allocated_storage
  max_allocated_storage   = var.max_allocated_storage
  storage_encrypted       = var.storage_encrypted
  kms_key_id              = var.kms_key_id
  storage_type            = "gp2"
  
  # High availability
  multi_az                = var.multi_az
  backup_retention_period = var.backup_retention_period
  backup_window           = var.backup_window
  maintenance_window      = var.maintenance_window
  
  # Performance
  performance_insights_enabled          = var.performance_insights_enabled
  performance_insights_kms_key_id       = var.kms_key_id
  monitoring_interval                   = var.monitoring_interval
  monitoring_role_arn                   = aws_iam_role.rds_enhanced_monitoring.arn
  enabled_cloudwatch_logs_exports       = var.enabled_cloudwatch_logs_exports
  
  # Network
  db_subnet_group_name    = var.db_subnet_group_name
  vpc_security_group_ids  = [aws_security_group.rds.id]
  
  # Security
  deletion_protection     = var.deletion_protection
  skip_final_snapshot     = var.skip_final_snapshot
  final_snapshot_identifier = var.final_snapshot_identifier
  
  # Parameter group
  parameter_group_name    = aws_db_parameter_group.main.name
  
  tags = merge(
    var.tags,
    {
      Name = var.db_instance_identifier != null ? var.db_instance_identifier : "${var.project_name}-${var.environment}-db"
    }
  )
  
  depends_on = [
    aws_iam_role_policy_attachment.rds_enhanced_monitoring
  ]
}

# Security Group for RDS
resource "aws_security_group" "rds" {
  name        = "${var.project_name}-${var.environment}-rds-sg"
  description = "Security group for RDS instance"
  vpc_id      = var.vpc_id

  ingress {
    description     = "PostgreSQL access from allowed security groups"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = var.allowed_security_group_ids
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-rds-sg"
    }
  )
}

# DB Parameter Group
resource "aws_db_parameter_group" "main" {
  name        = "${var.project_name}-${var.environment}-pg"
  family      = "postgres${split(".", var.engine_version)[0]}"
  description = "Parameter group for ${var.project_name} ${var.environment}"

  parameter {
    name  = "log_statement"
    value = "all"
  }

  parameter {
    name  = "log_min_duration_statement"
    value = "5000"
  }

  parameter {
    name  = "shared_preload_libraries"
    value = "pg_stat_statements"
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-pg"
    }
  )
}

# IAM Role for Enhanced Monitoring
resource "aws_iam_role" "rds_enhanced_monitoring" {
  name = "${var.project_name}-${var.environment}-rds-enhanced-monitoring-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "monitoring.rds.amazonaws.com"
        }
      }
    ]
  })

  tags = var.tags
}

resource "aws_iam_role_policy_attachment" "rds_enhanced_monitoring" {
  role       = aws_iam_role.rds_enhanced_monitoring.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole"
}
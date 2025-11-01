# ElastiCache Redis Module

resource "aws_elasticache_replication_group" "main" {
  replication_group_id          = "${var.project_name}-${var.environment}-redis"
  description = "Redis replication group for ${var.project_name} ${var.environment}"
  node_type                     = var.node_type
  engine                        = "redis"
  engine_version                = var.engine_version
  port                          = 6379
  parameter_group_name          = aws_elasticache_parameter_group.main.name
  subnet_group_name             = aws_elasticache_subnet_group.main.name
  security_group_ids            = [aws_security_group.elasticache.id]
  
  # High availability
  num_cache_clusters            = var.num_cache_nodes
  automatic_failover_enabled    = var.automatic_failover_enabled
  multi_az_enabled              = var.multi_az_enabled
  
  # Backup
  snapshot_retention_limit      = var.snapshot_retention_limit
  snapshot_window               = var.snapshot_window
  
  # Encryption
  at_rest_encryption_enabled    = var.at_rest_encryption_enabled
  transit_encryption_enabled    = var.transit_encryption_enabled
  auth_token                    = var.auth_token_enabled ? random_password.redis_auth_token[0].result : null
  kms_key_id                    = var.kms_key_id
  
  # Maintenance
  maintenance_window            = var.maintenance_window
  apply_immediately             = true
  
  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-redis"
    }
  )
}

# ElastiCache Subnet Group
resource "aws_elasticache_subnet_group" "main" {
  name        = "${var.project_name}-${var.environment}-redis-subnet-group"
  description = "Subnet group for Redis cluster"
  subnet_ids  = var.subnet_ids

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-redis-subnet-group"
    }
  )
}

# Security Group for ElastiCache
resource "aws_security_group" "elasticache" {
  name        = "${var.project_name}-${var.environment}-elasticache-sg"
  description = "Security group for ElastiCache Redis"
  vpc_id      = var.vpc_id

  ingress {
    description     = "Redis access from allowed security groups"
    from_port       = 6379
    to_port         = 6379
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
      Name = "${var.project_name}-${var.environment}-elasticache-sg"
    }
  )
}

# ElastiCache Parameter Group
resource "aws_elasticache_parameter_group" "main" {
  name        = "${var.project_name}-${var.environment}-redis-pg"
  family      = var.parameter_group_family
  description = "Parameter group for Redis cluster"

  parameter {
    name  = "tcp-keepalive"
    value = "0"
  }

  parameter {
    name  = "notify-keyspace-events"
    value = "lK"
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-redis-pg"
    }
  )
}

# Random password for Redis auth token
resource "random_password" "redis_auth_token" {
  count   = var.auth_token_enabled ? 1 : 0
  length  = 32
  special = false
}

# Secrets Manager secret for Redis auth token
resource "aws_secretsmanager_secret" "redis_auth_token" {
  count                   = var.auth_token_enabled ? 1 : 0
  name                    = "${var.project_name}/${var.environment}/redis/auth-token"
  recovery_window_in_days = 0

  tags = var.tags
}

resource "aws_secretsmanager_secret_version" "redis_auth_token" {
  count         = var.auth_token_enabled ? 1 : 0
  secret_id     = aws_secretsmanager_secret.redis_auth_token[0].id
  secret_string = random_password.redis_auth_token[0].result
}
# Outputs for ElastiCache module

output "cluster_id" {
  description = "ElastiCache cluster ID"
  value       = aws_elasticache_replication_group.main.id
}

output "cluster_arn" {
  description = "ElastiCache cluster ARN"
  value       = aws_elasticache_replication_group.main.arn
}

output "configuration_endpoint" {
  description = "ElastiCache configuration endpoint"
  value       = aws_elasticache_replication_group.main.configuration_endpoint_address
  sensitive   = true
}

output "primary_endpoint" {
  description = "ElastiCache primary endpoint"
  value       = aws_elasticache_replication_group.main.primary_endpoint_address
  sensitive   = true
}

output "reader_endpoint" {
  description = "ElastiCache reader endpoint"
  value       = aws_elasticache_replication_group.main.reader_endpoint_address
  sensitive   = true
}

output "port" {
  description = "ElastiCache port"
  value       = aws_elasticache_replication_group.main.port
}

output "member_clusters" {
  description = "ElastiCache member clusters"
  value       = aws_elasticache_replication_group.main.member_clusters
}

output "cluster_enabled" {
  description = "Indicates if cluster mode is enabled"
  value       = aws_elasticache_replication_group.main.cluster_enabled
}

output "engine_version_actual" {
  description = "Actual Redis engine version"
  value       = aws_elasticache_replication_group.main.engine_version_actual
}

output "elasticache_security_group_id" {
  description = "ElastiCache security group ID"
  value       = aws_security_group.elasticache.id
}

output "auth_token" {
  description = "Redis auth token"
  value       = var.auth_token_enabled ? random_password.redis_auth_token[0].result : null
  sensitive   = true
}

output "auth_token_secret_arn" {
  description = "ARN of Secrets Manager secret containing auth token"
  value       = var.auth_token_enabled ? aws_secretsmanager_secret.redis_auth_token[0].arn : null
  sensitive   = true
}
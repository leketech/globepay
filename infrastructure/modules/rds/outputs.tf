# Outputs for RDS module

output "db_instance_id" {
  description = "RDS instance ID"
  value       = aws_db_instance.main.id
}

output "db_instance_arn" {
  description = "RDS instance ARN"
  value       = aws_db_instance.main.arn
}

output "db_instance_endpoint" {
  description = "RDS instance endpoint"
  value       = aws_db_instance.main.endpoint
  sensitive   = true
}

output "db_instance_name" {
  description = "RDS instance name"
  value       = aws_db_instance.main.db_name
}

output "db_instance_username" {
  description = "RDS instance master username"
  value       = aws_db_instance.main.username
  sensitive   = true
}

output "db_instance_port" {
  description = "RDS instance port"
  value       = aws_db_instance.main.port
}

output "db_instance_status" {
  description = "RDS instance status"
  value       = aws_db_instance.main.status
}

output "db_instance_resource_id" {
  description = "RDS instance resource ID"
  value       = aws_db_instance.main.resource_id
}

output "db_security_group_id" {
  description = "RDS security group ID"
  value       = aws_security_group.rds.id
}

output "db_parameter_group_id" {
  description = "RDS parameter group ID"
  value       = aws_db_parameter_group.main.id
}
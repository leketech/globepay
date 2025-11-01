# Outputs for monitoring module

output "sns_topic_arn" {
  description = "SNS topic ARN"
  value       = aws_sns_topic.alerts.arn
}

output "sns_topic_name" {
  description = "SNS topic name"
  value       = aws_sns_topic.alerts.name
}

output "cloudwatch_alarms" {
  description = "CloudWatch alarm ARNs"
  value = {
    eks_cpu_utilization          = aws_cloudwatch_metric_alarm.eks_cpu_utilization.arn
    rds_cpu_utilization          = aws_cloudwatch_metric_alarm.rds_cpu_utilization.arn
    rds_free_storage_space       = aws_cloudwatch_metric_alarm.rds_free_storage_space.arn
    elasticache_cpu_utilization  = aws_cloudwatch_metric_alarm.elasticache_cpu_utilization.arn
    elasticache_swap_usage       = aws_cloudwatch_metric_alarm.elasticache_swap_usage.arn
  }
}
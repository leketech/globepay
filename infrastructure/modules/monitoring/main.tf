# Monitoring Module

# SNS Topic for Alerts
resource "aws_sns_topic" "alerts" {
  name = "${var.project_name}-${var.environment}-alerts"

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-alerts"
    }
  )
}

resource "aws_sns_topic_subscription" "email" {
  count     = var.alert_email != null ? 1 : 0
  topic_arn = aws_sns_topic.alerts.arn
  protocol  = "email"
  endpoint  = var.alert_email
}

# CloudWatch Alarms for EKS
resource "aws_cloudwatch_metric_alarm" "eks_cpu_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-eks-cpu-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EKS"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EKS cluster CPU utilization"
  alarm_actions       = [aws_sns_topic.alerts.arn]

  dimensions = {
    ClusterName = var.eks_cluster_name
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-eks-cpu-utilization"
    }
  )
}

# CloudWatch Alarms for RDS
resource "aws_cloudwatch_metric_alarm" "rds_cpu_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-rds-cpu-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors RDS instance CPU utilization"
  alarm_actions       = [aws_sns_topic.alerts.arn]

  dimensions = {
    DBInstanceIdentifier = var.rds_instance_id
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-rds-cpu-utilization"
    }
  )
}

resource "aws_cloudwatch_metric_alarm" "rds_free_storage_space" {
  alarm_name          = "${var.project_name}-${var.environment}-rds-free-storage-space"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "FreeStorageSpace"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "1073741824" # 1 GB in bytes
  alarm_description   = "This metric monitors RDS instance free storage space"
  alarm_actions       = [aws_sns_topic.alerts.arn]

  dimensions = {
    DBInstanceIdentifier = var.rds_instance_id
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-rds-free-storage-space"
    }
  )
}

# CloudWatch Alarms for ElastiCache
resource "aws_cloudwatch_metric_alarm" "elasticache_cpu_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-elasticache-cpu-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ElastiCache"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors ElastiCache cluster CPU utilization"
  alarm_actions       = [aws_sns_topic.alerts.arn]

  dimensions = {
    CacheClusterId = var.elasticache_cluster_id
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-elasticache-cpu-utilization"
    }
  )
}

resource "aws_cloudwatch_metric_alarm" "elasticache_swap_usage" {
  alarm_name          = "${var.project_name}-${var.environment}-elasticache-swap-usage"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "SwapUsage"
  namespace           = "AWS/ElastiCache"
  period              = "300"
  statistic           = "Average"
  threshold           = "52428800" # 50 MB in bytes
  alarm_description   = "This metric monitors ElastiCache cluster swap usage"
  alarm_actions       = [aws_sns_topic.alerts.arn]

  dimensions = {
    CacheClusterId = var.elasticache_cluster_id
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.project_name}-${var.environment}-elasticache-swap-usage"
    }
  )
}
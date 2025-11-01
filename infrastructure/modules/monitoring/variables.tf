# Variables for monitoring module

variable "project_name" {
  description = "Project name"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "eks_cluster_name" {
  description = "EKS cluster name"
  type        = string
  default     = null
}

variable "rds_instance_id" {
  description = "RDS instance ID"
  type        = string
  default     = null
}

variable "elasticache_cluster_id" {
  description = "ElastiCache cluster ID"
  type        = string
  default     = null
}

variable "alert_email" {
  description = "Email address for alerts"
  type        = string
  default     = null
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}
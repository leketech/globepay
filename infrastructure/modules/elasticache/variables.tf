# Variables for ElastiCache module

variable "project_name" {
  description = "Project name"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "subnet_ids" {
  description = "Subnet IDs"
  type        = list(string)
}

variable "allowed_security_group_ids" {
  description = "List of security group IDs allowed to access ElastiCache"
  type        = list(string)
}

# Redis configuration
variable "engine_version" {
  description = "Redis engine version"
  type        = string
  default     = "7.0"
}

variable "node_type" {
  description = "Cache node type"
  type        = string
  default     = "cache.t4g.micro"
}

variable "num_cache_nodes" {
  description = "Number of cache nodes"
  type        = number
  default     = 1
}

variable "parameter_group_family" {
  description = "Parameter group family"
  type        = string
  default     = "redis7"
}

# High availability
variable "automatic_failover_enabled" {
  description = "Enable automatic failover"
  type        = bool
  default     = false
}

variable "multi_az_enabled" {
  description = "Enable Multi-AZ"
  type        = bool
  default     = false
}

# Backup
variable "snapshot_retention_limit" {
  description = "Snapshot retention limit in days"
  type        = number
  default     = 7
}

variable "snapshot_window" {
  description = "Daily time range during which ElastiCache takes a snapshot"
  type        = string
  default     = "03:00-04:00"
}

# Encryption
variable "at_rest_encryption_enabled" {
  description = "Enable encryption at rest"
  type        = bool
  default     = true
}

variable "transit_encryption_enabled" {
  description = "Enable encryption in transit"
  type        = bool
  default     = true
}

variable "auth_token_enabled" {
  description = "Enable auth token"
  type        = bool
  default     = true
}

variable "kms_key_id" {
  description = "KMS key ID for encryption"
  type        = string
  default     = null
}

# Maintenance
variable "maintenance_window" {
  description = "Weekly time range during which system maintenance can occur"
  type        = string
  default     = "sun:05:00-sun:06:00"
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}
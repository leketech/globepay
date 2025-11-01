# Variables for S3 module

variable "project_name" {
  description = "Project name"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "buckets" {
  description = "Map of bucket configurations"
  type = map(object({
    versioning_enabled = bool
    lifecycle_rules = list(object({
      id      = string
      enabled = bool
      transition = optional(object({
        days          = number
        storage_class = string
      }))
      expiration = optional(object({
        days = number
      }))
    }))
  }))
  default = {
    assets = {
      versioning_enabled = true
      lifecycle_rules    = []
    }
  }
}

variable "kms_key_id" {
  description = "KMS key ID for encryption"
  type        = string
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}
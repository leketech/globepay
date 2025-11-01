# Variables for SES module

variable "project_name" {
  description = "Project name"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "domain_name" {
  description = "Domain name"
  type        = string
  default     = null
}

variable "route53_zone_id" {
  description = "Route53 zone ID"
  type        = string
  default     = null
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}

variable "verify_domain" {
  description = "Whether to verify the domain identity"
  type        = bool
  default     = true
}
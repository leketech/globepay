# Variables for production environment

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "globepay"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "prod"
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "availability_zones" {
  description = "Availability zones"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

variable "vpc_cidr" {
  description = "VPC CIDR block"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "Public subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}

variable "private_subnet_cidrs" {
  description = "Private subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.11.0/24", "10.0.12.0/24"]
}

variable "database_subnet_cidrs" {
  description = "Database subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.20.0/24", "10.0.21.0/24", "10.0.22.0/24"]
}

variable "eks_cluster_version" {
  description = "EKS cluster version"
  type        = string
  default     = "1.28"
}

variable "cluster_endpoint_public_access" {
  description = "Enable public access to EKS cluster endpoint"
  type        = bool
  default     = false
}

variable "cluster_endpoint_public_access_cidrs" {
  description = "List of CIDR blocks which can access the EKS cluster endpoint publicly"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "database_name" {
  description = "Database name"
  type        = string
  default     = "globepay_prod"
}

variable "database_username" {
  description = "Database username"
  type        = string
  default     = "globepay"
}

variable "database_password" {
  description = "Database password"
  type        = string
  default     = "your-secure-password-here"
}

variable "domain_name" {
  description = "Domain name"
  type        = string
  default     = "globepay.com"
}

variable "acm_certificate_arn" {
  description = "ACM certificate ARN"
  type        = string
}

variable "alert_email" {
  description = "Email address for alerts"
  type        = string
}

variable "route53_zone_id" {
  description = "Route53 zone ID"
  type        = string
  default     = null
}
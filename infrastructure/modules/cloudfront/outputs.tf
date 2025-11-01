# Outputs for CloudFront module

output "distribution_id" {
  description = "CloudFront distribution ID"
  value       = aws_cloudfront_distribution.main.id
}

output "distribution_arn" {
  description = "CloudFront distribution ARN"
  value       = aws_cloudfront_distribution.main.arn
}

output "distribution_domain_name" {
  description = "CloudFront distribution domain name"
  value       = aws_cloudfront_distribution.main.domain_name
}

output "distribution_etag" {
  description = "CloudFront distribution ETag"
  value       = aws_cloudfront_distribution.main.etag
}

output "origin_access_identity" {
  description = "CloudFront origin access identity"
  value       = aws_cloudfront_origin_access_identity.main.id
}

output "origin_access_identity_path" {
  description = "CloudFront origin access identity path"
  value       = aws_cloudfront_origin_access_identity.main.cloudfront_access_identity_path
}

output "origin_access_identity_caller_reference" {
  description = "CloudFront origin access identity caller reference"
  value       = aws_cloudfront_origin_access_identity.main.caller_reference
}
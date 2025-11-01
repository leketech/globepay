# Outputs for S3 module

output "bucket_ids" {
  description = "Map of bucket IDs"
  value       = { for k, bucket in aws_s3_bucket.buckets : k => bucket.id }
}

output "bucket_names" {
  description = "Map of bucket names"
  value       = { for k, bucket in aws_s3_bucket.buckets : k => bucket.bucket }
}

output "bucket_arns" {
  description = "Map of bucket ARNs"
  value       = { for k, bucket in aws_s3_bucket.buckets : k => bucket.arn }
}

output "bucket_regions" {
  description = "Map of bucket regions"
  value       = { for k, bucket in aws_s3_bucket.buckets : k => bucket.region }
}

output "bucket_domain_names" {
  description = "Map of bucket domain names"
  value       = { for k, bucket in aws_s3_bucket.buckets : k => bucket.bucket_regional_domain_name }
}
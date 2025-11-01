# Outputs for SES module

output "domain_identity_arn" {
  description = "SES domain identity ARN"
  value       = var.domain_name != null ? aws_ses_domain_identity.main[0].arn : null
}

output "domain_identity_verification_token" {
  description = "SES domain identity verification token"
  value       = var.domain_name != null ? aws_ses_domain_identity.main[0].verification_token : null
  sensitive   = true
}

output "dkim_tokens" {
  description = "SES DKIM tokens"
  value       = var.domain_name != null ? aws_ses_domain_dkim.main[0].dkim_tokens : null
  sensitive   = true
}

output "configuration_set_name" {
  description = "SES configuration set name"
  value       = aws_ses_configuration_set.main.name
}

output "receipt_rule_set_name" {
  description = "SES receipt rule set name"
  value       = aws_ses_receipt_rule_set.main.rule_set_name
}
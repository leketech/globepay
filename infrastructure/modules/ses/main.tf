# SES Module

resource "aws_ses_domain_identity" "main" {
  count  = var.domain_name != null ? 1 : 0
  domain = var.domain_name
}

resource "aws_ses_domain_identity_verification" "main" {
  count  = var.verify_domain && var.domain_name != null ? 1 : 0
  domain = aws_ses_domain_identity.main[0].domain

  depends_on = [aws_route53_record.ses_verification]
}

resource "aws_ses_domain_dkim" "main" {
  count  = var.domain_name != null ? 1 : 0
  domain = aws_ses_domain_identity.main[0].domain
}

resource "aws_route53_record" "ses_verification" {
  count   = var.route53_zone_id != null && var.domain_name != null ? 1 : 0
  zone_id = var.route53_zone_id
  name    = "_amazonses.${var.domain_name}"
  type    = "TXT"
  ttl     = "600"
  records = [aws_ses_domain_identity.main[0].verification_token]
}

resource "aws_route53_record" "ses_dkim" {
  count   = var.route53_zone_id != null && var.domain_name != null ? 3 : 0
  zone_id = var.route53_zone_id
  name    = "${element(aws_ses_domain_dkim.main[0].dkim_tokens, count.index)}._domainkey.${var.domain_name}"
  type    = "CNAME"
  ttl     = "600"
  records = ["${element(aws_ses_domain_dkim.main[0].dkim_tokens, count.index)}.dkim.amazonses.com"]
}

resource "aws_route53_record" "ses_mx" {
  count   = var.route53_zone_id != null && var.domain_name != null ? 1 : 0
  zone_id = var.route53_zone_id
  name    = var.domain_name
  type    = "MX"
  ttl     = "600"
  records = ["10 inbound-smtp.${var.aws_region}.amazonaws.com"]
}

resource "aws_ses_receipt_rule_set" "main" {
  rule_set_name = "${var.project_name}-${var.environment}-receipt-rule-set"
}

resource "aws_ses_active_receipt_rule_set" "main" {
  count         = var.domain_name != null ? 1 : 0
  rule_set_name = aws_ses_receipt_rule_set.main.rule_set_name
}

resource "aws_ses_configuration_set" "main" {
  name = "${var.project_name}-${var.environment}-config-set"
}

resource "aws_ses_event_destination" "main" {
  name                   = "${var.project_name}-${var.environment}-event-destination"
  configuration_set_name = aws_ses_configuration_set.main.name

  cloudwatch_destination {
    default_value  = "default"
    dimension_name = "dimension"
    value_source   = "emailHeader"
  }

  matching_types = ["bounce", "complaint", "delivery", "send", "reject"]
}
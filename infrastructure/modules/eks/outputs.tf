# Outputs for EKS module

output "cluster_name" {
  description = "EKS cluster name"
  value       = aws_eks_cluster.main.name
}

output "cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = aws_eks_cluster.main.endpoint
  sensitive   = true
}

output "cluster_certificate_authority_data" {
  description = "EKS cluster certificate authority data"
  value       = aws_eks_cluster.main.certificate_authority[0].data
  sensitive   = true
}

output "cluster_id" {
  description = "EKS cluster ID"
  value       = aws_eks_cluster.main.id
}

output "cluster_arn" {
  description = "EKS cluster ARN"
  value       = aws_eks_cluster.main.arn
}

output "cluster_oidc_issuer_url" {
  description = "EKS cluster OIDC issuer URL"
  value       = aws_eks_cluster.main.identity[0].oidc[0].issuer
}

output "cluster_version" {
  description = "EKS cluster version"
  value       = aws_eks_cluster.main.version
}

output "cluster_security_group_id" {
  description = "EKS cluster security group ID"
  value       = aws_eks_cluster.main.vpc_config[0].cluster_security_group_id
}

output "node_group_role_arn" {
  description = "EKS node group role ARN"
  value       = aws_iam_role.node_group.arn
}

output "cluster_primary_security_group_id" {
  description = "EKS cluster primary security group ID"
  value       = aws_eks_cluster.main.vpc_config != null && length(aws_eks_cluster.main.vpc_config) > 0 && aws_eks_cluster.main.vpc_config[0].security_group_ids != null && length(aws_eks_cluster.main.vpc_config[0].security_group_ids) > 0 ? tolist(aws_eks_cluster.main.vpc_config[0].security_group_ids)[0] : ""
}
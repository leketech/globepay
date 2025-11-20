# Deployment Workflow Updates

This document explains the changes made to the deployment workflows and ArgoCD configurations to improve the CI/CD and GitOps integration.

## Summary of Changes

1. **Updated Deployment Workflow**:
   - Modified [deploy-prod.yml](file:///c%3A/Users/Leke/payment/globepay/.github/workflows/deploy-prod.yml) to use Kustomize instead of Helm for deployments
   - Removed all Helm-related configurations and dependencies

2. **Updated ArgoCD Application Configurations**:
   - Verified that no hardcoded image tag parameters exist in ArgoCD application specs
   - Maintained clean separation between CI/CD pipeline and GitOps workflow
   - Updated applications to use Kustomize paths instead of Helm charts

3. **Added Verification Scripts**:
   - Created [verify-argocd-status.sh](file:///c%3A/Users/Leke/payment/globepay/scripts/verify-argocd-status.sh) to check application health and sync status
   - Integrated verification steps into the deployment workflow

## Current Workflow

The current workflow uses Kustomize for deployments. This approach:
- Uses Kustomize overlays to manage environment-specific configurations
- Updates image tags directly in the Kustomize overlay files during deployment
- Applies configurations using `kubectl apply -k` command
- Provides immediate deployment feedback

## How to Use

1. Ensure the workflow file [.github/workflows/deploy-prod.yml](file:///c%3A/Users/Leke/payment/globepay/.github/workflows/deploy-prod.yml) is in place
2. Trigger deployment by pushing to the `main` branch
3. The workflow will automatically deploy using Kustomize

## Verification Process

The workflow includes steps to verify:
1. Kubernetes deployment rollout status
2. ArgoCD application health and sync status
3. Service accessibility through smoke tests

The verification script [scripts/verify-argocd-status.sh](file:///c%3A/Users/Leke/payment/globepay/scripts/verify-argocd-status.sh) checks that:
- Applications are in "Healthy" state
- Applications are in "Synced" state
- All required applications are present in ArgoCD

## Recommendations

1. **Use the current Kustomize-based workflow** as it's simpler and doesn't require Helm dependencies
2. **Remove any unused Helm-related files** to keep the repository clean
3. **Regularly verify ArgoCD status** after deployments to ensure applications are healthy and synced

## Troubleshooting

If deployments fail:
1. Check ArgoCD UI for detailed sync and health information
2. Verify that the EKS cluster is accessible
3. Ensure all required secrets are configured in GitHub
4. Check the logs of failed workflow steps for specific error messages
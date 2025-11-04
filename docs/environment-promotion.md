# Environment Promotion Process

This document explains how to promote changes from development to production environments in the Globepay system.

## Overview

The Globepay system uses a multi-environment deployment strategy:
- **Development (dev)**: Initial testing environment
- **Staging (staging)**: Pre-production testing environment
- **Production (prod)**: Live user-facing environment

Changes flow from dev → staging → prod following a promotion process that ensures quality and consistency.

## Automated Promotion

### Using Make Commands

You can promote changes using the Makefile targets:

```bash
# Deploy to all environments (dev -> staging -> prod)
make deploy-all

# Deploy to specific environments
make deploy-dev
make deploy-staging
make deploy-prod
```

### Using the Promotion Script Directly

The [scripts/promote-changes.sh](file:///c%3A/Users/Leke/Globepay/globepay/scripts/promote-changes.sh) script provides more granular control:

```bash
# Promote through all environments
./scripts/promote-changes.sh all

# Promote to specific environment only
./scripts/promote-changes.sh dev
./scripts/promote-changes.sh staging
./scripts/promote-changes.sh prod
```

## Manual Promotion Process

For manual deployments, use kubectl with kustomize:

```bash
# Deploy to development
kubectl apply -k k8s/overlays/dev

# Deploy to staging
kubectl apply -k k8s/overlays/staging

# Deploy to production
kubectl apply -k k8s/overlays/prod
```

## GitHub Actions Workflow

Changes can also be promoted through the GitHub Actions workflow:

1. Go to the "Actions" tab in the repository
2. Select "Promote Changes Across Environments"
3. Click "Run workflow"
4. Choose the target environment
5. Click "Run workflow"

## Environment-Specific Configurations

Each environment has its own kustomization configuration in [k8s/overlays/](file:///c%3A/Users/Leke/Globepay/globepay/k8s/overlays/):

- [k8s/overlays/dev/kustomization.yaml](file:///c%3A/Users/Leke/Globepay/globepay/k8s/overlays/dev/kustomization.yaml): Development configuration
- [k8s/overlays/staging/kustomization.yaml](file:///c%3A/Users/Leke/Globepay/globepay/k8s/overlays/staging/kustomization.yaml): Staging configuration
- [k8s/overlays/prod/kustomization.yaml](file:///c%3A/Users/Leke/Globepay/globepay/k8s/overlays/prod/kustomization.yaml): Production configuration

## Best Practices

1. **Always test in dev first**: Deploy and verify changes in development before promoting
2. **Run tests between promotions**: Ensure automated tests pass before promoting to the next environment
3. **Monitor after deployment**: Check logs and metrics after each deployment
4. **Rollback if needed**: Use kubectl to rollback if issues are discovered

## Rollback Process

If issues are discovered after promotion, you can rollback using:

```bash
# Rollback deployment in specific environment
kubectl rollout undo deployment/backend -n globepay-[environment]
kubectl rollout undo deployment/frontend -n globepay-[environment]
```
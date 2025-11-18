# Helm Chart for Globepay

This directory contains the Helm chart for deploying the Globepay application to Kubernetes.

## Prerequisites

- Kubernetes 1.16+
- Helm 3.0+

## Chart Structure

```
globepay/
├── Chart.yaml          # Chart information
├── values.yaml         # Default configuration values
├── values-prod.yaml    # Production configuration values
├── templates/          # Kubernetes manifest templates
│   ├── _helpers.tpl    # Helper templates
│   ├── backend-deployment.yaml
│   ├── backend-service.yaml
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   ├── database-deployment.yaml
│   ├── database-service.yaml
│   ├── redis-deployment.yaml
│   ├── redis-service.yaml
│   ├── ingress.yaml
│   ├── secrets.yaml
│   ├── configmap.yaml
│   └── migration-job.yaml
```

## Configuration

The following table lists the configurable parameters of the Globepay chart and their default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.environment` | Deployment environment | `production` |
| `global.namespace` | Kubernetes namespace | `globepay-prod` |
| `backend.replicaCount` | Number of backend replicas | `3` |
| `backend.image.repository` | Backend image repository | `907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend` |
| `backend.image.tag` | Backend image tag | `latest` |
| `backend.image.pullPolicy` | Backend image pull policy | `IfNotPresent` |
| `frontend.replicaCount` | Number of frontend replicas | `3` |
| `frontend.image.repository` | Frontend image repository | `907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-frontend` |
| `frontend.image.tag` | Frontend image tag | `latest` |
| `frontend.image.pullPolicy` | Frontend image pull policy | `IfNotPresent` |

## Deploying to Production

To deploy the application to production using Helm:

```bash
helm upgrade --install globepay-prod ./helm/globepay \
  -f ./helm/globepay/values-prod.yaml \
  --namespace globepay-prod \
  --create-namespace \
  --set secrets.dbPassword="your-db-password" \
  --set secrets.jwtSecret="your-jwt-secret"
```

## GitHub Actions Integration

The deployment is automated through GitHub Actions using the workflow defined in `.github/workflows/deploy-prod-helm.yml`.
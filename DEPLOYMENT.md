# Globepay Deployment Guide

This guide provides detailed instructions for deploying Globepay to various environments.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Local Development Deployment](#local-development-deployment)
- [Staging Environment Deployment](#staging-environment-deployment)
- [Production Environment Deployment](#production-environment-deployment)
- [Infrastructure Deployment](#infrastructure-deployment)
- [Application Deployment](#application-deployment)
- [Monitoring Setup](#monitoring-setup)
- [Rollback Procedures](#rollback-procedures)
- [Troubleshooting](#troubleshooting)

## Prerequisites

Before deploying Globepay, ensure you have the following tools installed:

1. Docker (24.0+)
2. Docker Compose (2.20+)
3. kubectl (1.28+)
4. Helm (3.12+)
5. Terraform (1.5+)
6. AWS CLI (2.13+)
7. ArgoCD CLI (optional but recommended)

## Local Development Deployment

For local development, we use Docker Compose to run all services:

```bash
# Start the development environment
make dev-up

# View logs
make dev-logs

# Stop the development environment
make dev-down
```

This will start:
- PostgreSQL database
- Redis cache
- Backend API
- Frontend application
- Prometheus monitoring
- Grafana dashboard

Access the services at:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Grafana: http://localhost:3001
- Prometheus: http://localhost:9091

## Staging Environment Deployment

### 1. Docker Compose Deployment (Recommended)

For a simpler staging deployment using Docker Compose:

```bash
# Option 1: Using the deployment script
./scripts/deploy-staging.sh

# Option 2: Using Makefile target
make deploy-staging-docker
```

This will:
1. Load environment variables from `.env.staging`
2. Build and start all services using `docker-compose.staging.yml`

### 2. Infrastructure Setup

```bash
# Navigate to staging environment
cd infrastructure/environments/staging

# Initialize Terraform
terraform init

# Review the plan
terraform plan

# Apply the configuration
terraform apply
```

### 3. Application Deployment

```bash
# Build and push Docker images
docker build -t globepay-backend:staging ./backend
docker build -t globepay-frontend:staging ./frontend

# Deploy using ArgoCD
kubectl apply -f k8s/argocd/application-staging.yaml
```

## Production Environment Deployment

### 1. Docker Compose Deployment (For Testing Only)

For testing purposes, you can deploy to a production-like environment using Docker Compose:

```bash
# Option 1: Using the deployment script
./scripts/deploy-prod.sh

# Option 2: Using Makefile target
make deploy-prod-docker
```

> ⚠️ **Warning**: This Docker Compose deployment is for testing purposes only. For actual production deployments, use the Kubernetes deployment method described below.

### 2. Infrastructure Deployment

#### Step 1: Setup Terraform Backend

You can either run the commands manually or use the provided setup script:

```bash
# Option 1: Run the setup script (Linux/Mac)
cd infrastructure
./scripts/setup-terraform-backend.sh

# Option 2: Run the setup script (Windows)
cd infrastructure
./scripts/setup-terraform-backend.ps1

# Option 3: Run commands manually
# Create S3 bucket for state storage
aws s3api create-bucket \
  --bucket globepay-terraform-state-prod \
  --region us-east-1

# Enable versioning
aws s3api put-bucket-versioning \
  --bucket globepay-terraform-state-prod \
  --versioning-configuration Status=Enabled

# Create DynamoDB table for state locking
aws dynamodb create-table \
  --table-name globepay-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --region us-east-1
```

#### Step 2: Deploy Infrastructure

```bash
# Navigate to production environment
cd infrastructure/environments/prod

# Initialize Terraform with remote backend
terraform init

# Review the plan
terraform plan

# Apply the configuration
terraform apply
```

### 2. Application Deployment

#### Step 1: Build and Push Docker Images

```bash
# Build backend image
docker build -t 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend:latest ./backend

# Build frontend image
docker build -t 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-frontend:latest ./frontend

# Push images to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 907849381252.dkr.ecr.us-east-1.amazonaws.com

docker push 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend:latest
docker push 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-frontend:latest
```

#### Step 2: Deploy with ArgoCD

```bash
# Update kubeconfig
aws eks update-kubeconfig --name globepay-prod-eks --region us-east-1

# Apply ArgoCD application
kubectl apply -f k8s/argocd/application-prod.yaml

# Monitor deployment
kubectl get applications -n argocd
```

## Infrastructure Deployment

### VPC and Networking

The infrastructure is deployed using Terraform modules:

1. **VPC Module**: Creates the virtual private cloud with public and private subnets
2. **EKS Module**: Sets up the Kubernetes cluster
3. **RDS Module**: Creates the PostgreSQL database
4. **ElastiCache Module**: Sets up Redis cache
5. **S3 Module**: Creates storage buckets
6. **CloudFront Module**: Configures CDN
7. **ALB Module**: Sets up application load balancer

### Database Migration

After infrastructure deployment, run database migrations:

```bash
# Port-forward to the migration pod
kubectl port-forward -n globepay-prod deploy/backend-migration 8080:8080

# Run migrations
kubectl exec -it -n globepay-prod deploy/backend-migration -- /app/migrate up
```

## Application Deployment

### Backend Deployment

The backend is deployed as a Kubernetes deployment with:

- Auto-scaling based on CPU and memory usage
- Health checks (liveness and readiness probes)
- Resource limits and requests
- Security context (non-root user)
- ConfigMaps for configuration
- Secrets for sensitive data

### Frontend Deployment

The frontend is deployed as a Kubernetes deployment with:

- Nginx as the web server
- Static asset caching
- Security headers
- Health checks
- Resource limits and requests

## Monitoring Setup

### Prometheus Configuration

Prometheus is deployed using the kube-prometheus-stack Helm chart:

```bash
# Add Helm repository
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install Prometheus stack
helm install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --values monitoring/prometheus/values.yaml
```

### Grafana Dashboards

Grafana dashboards are automatically provisioned:

1. **API Dashboard**: Shows API performance metrics
2. **Infrastructure Dashboard**: Displays Kubernetes and AWS resource usage
3. **Database Dashboard**: Monitors PostgreSQL performance
4. **Business Metrics Dashboard**: Tracks transaction volumes and user growth

### Alerting Rules

Alerts are configured for:

- High error rates (> 5%)
- High latency (> 1s P95)
- High CPU usage (> 80%)
- High memory usage (> 90%)
- Database connection pool exhaustion
- Low cache hit rates (< 80%)

## Rollback Procedures

### Rolling Back Infrastructure Changes

```bash
# Navigate to the environment
cd infrastructure/environments/prod

# Review the state
terraform state list

# Rollback to a previous state
terraform apply -target=module.eks
```

### Rolling Back Application Deployments

```bash
# Rollback to previous version using ArgoCD
argocd app rollback globepay-backend-prod

# Or using kubectl
kubectl rollout undo deployment/backend -n globepay-prod
```

### Database Rollback

```bash
# Rollback last migration
make migration-down

# Or using kubectl
kubectl exec -it -n globepay-prod deploy/backend-migration -- /app/migrate down 1
```

## Troubleshooting

### Common Issues

1. **Database Connection Issues**
   - Check if the RDS instance is running
   - Verify security group rules
   - Confirm database credentials

2. **Application Not Starting**
   - Check pod logs: `kubectl logs -n globepay-prod deploy/backend`
   - Verify ConfigMaps and Secrets
   - Check resource limits

3. **Monitoring Not Working**
   - Verify Prometheus targets
   - Check service discovery
   - Confirm network policies

### Useful Commands

```bash
# Check pod status
kubectl get pods -n globepay-prod

# Check service status
kubectl get services -n globepay-prod

# Check logs
kubectl logs -f -n globepay-prod deploy/backend

# Describe pod for detailed info
kubectl describe pod -n globepay-prod <pod-name>

# Port forward for debugging
kubectl port-forward -n globepay-prod svc/backend 8080:80
```

### Health Checks

Monitor the following endpoints:

- Backend health: `GET /health`
- Backend readiness: `GET /ready`
- Frontend health: `GET /health`
- Database: `pg_isready`
- Redis: `redis-cli ping`

This deployment guide provides a comprehensive overview of deploying Globepay to various environments. Always test deployments in staging before applying to production.
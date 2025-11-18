# Globepay - Production-Ready Fintech Application

<div align="center">

![Globepay Logo](https://img.shields.io/badge/Globepay-Fintech-blue)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org)
[![React Version](https://img.shields.io/badge/React-18+-61DAFB?logo=react)](https://reactjs.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

**Send Money Borderless - Instantly transfer funds to over 190 countries**

[Features](#features) • [Architecture](#architecture) • [Quick Start](#quick-start) • [Deployment](#deployment) • [Documentation](#documentation)

</div>

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Technology Stack](#technology-stack)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)
- [Monitoring](#monitoring)
- [Security](#security)
- [Contributing](#contributing)

---

## 🎯 Overview

Globepay is a production-ready fintech application that enables seamless money transfers across 190+ countries. Built with modern cloud-native technologies, it provides a secure, scalable, and compliant platform for international money transfers.

### Key Highlights

- 🌍 **Global Coverage**: Support for 190+ countries
- ⚡ **Real-time Transfers**: Instant money transfer processing
- 🔒 **Bank-grade Security**: PCI-DSS compliance ready
- 📊 **Full Observability**: Complete monitoring with Prometheus & Grafana
- 🚀 **Auto-scaling**: Kubernetes-based infrastructure
- 🔄 **CI/CD Pipeline**: Automated deployments with GitHub Actions & ArgoCD

---

## ✨ Features

### Core Functionality

#### User Management
- ✅ Email & Social Authentication (OAuth2)
- ✅ KYC/AML Verification
- ✅ Multi-factor Authentication (2FA)
- ✅ User Profile Management

#### Money Transfers
- ✅ International Wire Transfers
- ✅ Real-time Exchange Rates
- ✅ Multiple Currency Support (50+ currencies)
- ✅ Transfer Status Tracking
- ✅ Scheduled Transfers
- ✅ Recipient Management

#### Transaction Management
- ✅ Complete Transaction History
- ✅ Advanced Search & Filtering
- ✅ PDF/CSV Export
- ✅ Receipt Generation
- ✅ Transaction Analytics

#### Compliance & Security
- ✅ AML Transaction Monitoring
- ✅ KYC Document Verification
- ✅ Fraud Detection System
- ✅ Regulatory Reporting
- ✅ Data Encryption (At-rest & In-transit)

#### Notifications
- ✅ Email Notifications (AWS SES)
- ✅ SMS Notifications (AWS SNS)
- ✅ In-app Notifications
- ✅ Push Notifications

---

## 🏗️ Architecture

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CloudFront (CDN)                         │
└─────────────────┬───────────────────────────────────────────────┘
                  │
       ┌──────────┴──────────┐
       │                     │
       ▼                     ▼
┌─────────────┐      ┌─────────────────┐
│   S3 Bucket │      │  AWS ALB/Ingress│
│  (Frontend) │      │   (API Gateway)  │
└─────────────┘      └────────┬─────────┘
                              │
                   ┌──────────┴──────────┐
                   │   EKS Cluster       │
                   │ ┌─────────────────┐ │
                   │ │  Backend Pods   │ │
                   │ │   (Auto-scale)  │ │
                   │ └────────┬────────┘ │
                   │          │          │
                   │ ┌────────┴────────┐ │
                   │ │  Worker Pods    │ │
                   │ └─────────────────┘ │
                   └──────────┬──────────┘
                              │
         ┌────────────────────┼────────────────────┐
         │                    │                    │
         ▼                    ▼                    ▼
  ┌─────────────┐      ┌─────────────┐     ┌─────────────┐
  │  RDS        │      │ ElastiCache │     │     SQS     │
  │ (Postgres)  │      │   (Redis)   │     │  (Queues)   │
  └─────────────┘      └─────────────┘     └─────────────┘
```

### Infrastructure Components

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Compute** | Amazon EKS | Container orchestration |
| **Database** | Amazon RDS (PostgreSQL) | Primary data store |
| **Cache** | Amazon ElastiCache (Redis) | Session & data caching |
| **Storage** | Amazon S3 | Document storage |
| **CDN** | CloudFront | Static asset delivery |
| **Queue** | Amazon SQS | Async job processing |
| **Email** | Amazon SES | Email notifications |
| **Secrets** | AWS Secrets Manager | Secrets management |
| **Monitoring** | Prometheus + Grafana | Metrics & visualization |
| **Logging** | Loki + Promtail | Log aggregation |
| **Tracing** | Jaeger | Distributed tracing |

---

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **ORM**: SQLx (Database interactions)
- **Authentication**: JWT + OAuth2
- **API Documentation**: Swagger/OpenAPI
- **Testing**: Go testing framework + Testify

### Frontend
- **Framework**: React 18+ with TypeScript
- **State Management**: Redux Toolkit
- **Styling**: Tailwind CSS
- **Build Tool**: Vite
- **Testing**: Jest + React Testing Library + Playwright

### Infrastructure
- **IaC**: Terraform
- **Container Orchestration**: Kubernetes (Amazon EKS)
- **CI/CD**: GitHub Actions + ArgoCD
- **Service Mesh**: Istio (Optional)

### Observability
- **Metrics**: Prometheus
- **Visualization**: Grafana
- **Logging**: Loki + Promtail
- **Tracing**: Jaeger / AWS X-Ray
- **APM**: OpenTelemetry

### Databases
- **Primary DB**: PostgreSQL 15 (Multi-AZ)
- **Cache**: Redis 7
- **Search**: (Optional) Elasticsearch

---

## 📦 Prerequisites

### Windows Users (WSL)

For Windows users, we recommend using WSL2 for development. See [WSL Development Guide](docs/wsl-development.md) for detailed instructions.

### Required Software

| Software | Version | Purpose |
|----------|---------|---------|
| Docker | 24.0+ | Container runtime |
| Docker Compose | 2.20+ | Local development |
| Go | 1.21+ | Backend development |
| Node.js | 20+ | Frontend development |
| Terraform | 1.5+ | Infrastructure provisioning |
| kubectl | 1.28+ | Kubernetes management |
| AWS CLI | 2.13+ | AWS resource management |
| Helm | 3.12+ | Kubernetes package manager |

### AWS Requirements

1. **AWS Account** with appropriate permissions
2. **IAM User/Role** with following policies:
   - AmazonEKSClusterPolicy
   - AmazonEKSServicePolicy
   - AmazonEC2FullAccess
   - AmazonRDSFullAccess
   - AmazonS3FullAccess
   - AmazonVPCFullAccess
3. **Route53 Hosted Zone** (for DNS)
4. **ACM Certificate** (for SSL/TLS)

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/your-org/globepay.git
cd globepay
```

### 2. Environment Setup

```bash
# Run the setup script
chmod +x scripts/setup-dev-environment.sh
./scripts/setup-dev-environment.sh

# Or use Make
make setup
```

### 3. Configure Environment Variables

```bash
# Backend configuration
cp backend/.env.example backend/.env
# Edit backend/.env with your configurations

# Frontend configuration
cp frontend/.env.example frontend/.env
# Edit frontend/.env with your configurations
```

### 4. Start Development Environment

```bash
# Start all services
make dev-up

# View logs
make dev-logs

# Stop services
make dev-down
```

### 5. Access Applications

| Service | URL | Credentials |
|---------|-----|-------------|
| Frontend | http://localhost:3000 | - |
| Backend API | http://localhost:8080 | - |
| API Docs | http://localhost:8080/swagger | - |
| Grafana | http://localhost:3001 | admin/admin |
| Prometheus | http://localhost:9091 | - |

---

## 💻 Development

### Project Structure

```
globepay/
├── backend/                 # Go backend application
│   ├── cmd/                # Application entrypoints
│   ├── internal/           # Private application code
│   │   ├── api/           # HTTP handlers & routes
│   │   ├── domain/        # Business logic & entities
│   │   ├── repository/    # Data access layer
│   │   ├── service/       # Business services
│   │   └── infrastructure/# External dependencies
│   └── test/              # Tests
├── frontend/               # React frontend application
│   ├── src/
│   │   ├── components/    # React components
│   │   ├── pages/         # Page components
│   │   ├── services/      # API services
│   │   ├── store/         # Redux store
│   │   └── hooks/         # Custom hooks
│   └── public/            # Static assets
├── infrastructure/         # Terraform IaC
│   ├── modules/           # Reusable modules
│   └── environments/      # Environment configs
├── k8s/                   # Kubernetes manifests
│   ├── base/              # Base configurations
│   └── overlays/          # Environment overlays
└── monitoring/            # Observability configs
```

### Backend Development

```bash
# Navigate to backend
cd backend

# Install dependencies
go mod download

# Run tests
go test ./...

# Run with hot reload
air

# Build binary
go build -o bin/api cmd/api/main.go

# Run linter
golangci-lint run
```

### Frontend Development

```bash
# Navigate to frontend
cd frontend

# Install dependencies
npm install

# Start dev server
npm run dev

# Run tests
npm test

# Run E2E tests
npm run test:e2e

# Build for production
npm run build

# Run linter
npm run lint
```

### Database Migrations

```bash
# Create new migration
make migration-create NAME=add_users_table

# Run migrations up
make migration-up

# Rollback last migration
make migration-down
```

---

## 🧪 Testing

### Backend Testing

```bash
# Run all tests
make test-backend

# Run with coverage
cd backend
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run integration tests
go test -v -tags=integration ./test/integration/...

# Run specific test
go test -v ./internal/service -run TestTransferService
```

### Frontend Testing

```bash
# Run unit tests
make test-frontend

# Run with coverage
cd frontend
npm run test:coverage

# Run E2E tests
npm run test:e2e

# Run E2E tests in UI mode
npm run test:e2e:ui
```

---

## 🚢 Deployment

### Infrastructure Deployment

#### Step 1: Initialize Terraform Backend

```bash
cd infrastructure

# Create S3 bucket for state
aws s3api create-bucket \
  --bucket globepay-terraform-state-prod \
  --region us-east-1

# Enable versioning
aws s3api put-bucket-versioning \
  --bucket globepay-terraform-state-prod \
  --versioning-configuration Status=Enabled

# Create DynamoDB table for locking
aws dynamodb create-table \
  --table-name globepay-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --region us-east-1
```

#### Step 2: Configure Environment

```bash
cd infrastructure/environments/prod

# Copy and edit terraform.tfvars
cp terraform.tfvars.example terraform.tfvars

# Edit with your values
vim terraform.tfvars
```

Example `terraform.tfvars`:

```hcl
project_name     = "globepay"
environment      = "prod"
aws_region       = "us-east-1"
vpc_cidr         = "10.0.0.0/16"
availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"]

# Database
database_name     = "globepay"
database_username = "globepay_admin"

# Domain
domain_name            = "globepay.com"
acm_certificate_arn    = "arn:aws:acm:us-east-1:xxx:certificate/xxx"

# Alerting
alert_email = "alerts@globepay.com"
```

#### Step 3: Deploy Infrastructure

```bash
# Initialize Terraform
terraform init

# Review plan
terraform plan

# Apply changes
terraform apply

# Save outputs
terraform output > outputs.txt
```

### Application Deployment

#### Step 1: Setup ECR Repositories

```bash
# Create ECR repositories
aws ecr create-repository --repository-name globepay-backend --region us-east-1
aws ecr create-repository --repository-name globepay-frontend --region us-east-1
```

#### Step 2: Configure kubectl

```bash
# Update kubeconfig
aws eks update-kubeconfig \
  --name globepay-prod-eks \
  --region us-east-1

# Verify connection
kubectl get nodes
```

#### Step 3: Install ArgoCD

```bash
# Create namespace
kubectl create namespace argocd

# Install ArgoCD
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Expose ArgoCD server
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

# Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d

# Access ArgoCD UI
ARGOCD_URL=$(kubectl get svc argocd-server -n argocd -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
echo "ArgoCD URL: https://${ARGOCD_URL}"
```

#### Step 4: Deploy Monitoring Stack

```bash
# Add Prometheus Helm repo
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install Prometheus
helm install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --values monitoring/prometheus/values.yaml

# Install Loki
helm repo add grafana https://grafana.github.io/helm-charts
helm install loki grafana/loki-stack \
  --namespace monitoring \
  --values monitoring/loki/values.yaml
```

#### Step 5: Configure Secrets

```bash
# Create namespace
kubectl create namespace globepay-prod

# Create secrets from AWS Secrets Manager
aws secretsmanager create-secret \
  --name globepay-prod-backend \
  --secret-string file://secrets.json \
  --region us-east-1

# Create Kubernetes secret
kubectl create secret generic backend-secrets \
  --from-literal=database-url="postgresql://..." \
  --from-literal=redis-url="redis://..." \
  --from-literal=jwt-secret="..." \
  --namespace globepay-prod
```

#### Step 6: Deploy Application with ArgoCD

```bash
# Apply ArgoCD project
kubectl apply -f k8s/argocd/project.yaml

# Apply application
kubectl apply -f k8s/argocd/application.yaml

# Watch deployment
kubectl get applications -n argocd
argocd app get globepay-backend-prod
```

### CI/CD Pipeline Setup

#### Step 1: Configure GitHub Secrets

Add the following secrets to your GitHub repository:
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_REGION`
- `DOCKERHUB_USERNAME` (optional)
- `DOCKERHUB_TOKEN` (optional)
- `SLACK_WEBHOOK`
- `PROD_API_URL`
- `CLOUDFRONT_DISTRIBUTION_ID`

#### Step 2: Enable GitHub Actions

The workflows in `.github/workflows/` will automatically run on:

- Push to `develop`: Deploy to Dev
- Push to `staging`: Deploy to Staging
- Push to `main`: Deploy to Production (requires manual approval)

#### Step 3: Manual Deployment

```bash
# Trigger deployment manually
gh workflow run deploy-prod.yml
```

### Environment Promotion Process

Changes are promoted from development to production using an automated process:

```bash
# Deploy to all environments (dev -> staging -> prod)
make deploy-all

# Deploy to specific environments
make deploy-dev
make deploy-staging
make deploy-prod
```

See [Environment Promotion Documentation](docs/environment-promotion.md) for detailed information.

---

## 📊 Monitoring

### Access Monitoring Tools

```bash
# Port-forward Grafana
kubectl port-forward -n monitoring svc/prometheus-grafana 3000:80

# Port-forward Prometheus
kubectl port-forward -n monitoring svc/prometheus-kube-prometheus-prometheus 9090:9090

# Access URLs
# Grafana: http://localhost:3000 (admin/prom-operator)
# Prometheus: http://localhost:9090
```

### Key Metrics Monitored

| Metric | Description | Alert Threshold |
|--------|-------------|-----------------|
| API Error Rate | 5xx errors per second | > 5% |
| API Latency (P95) | 95th percentile response time | > 1s |
| Pod CPU Usage | Container CPU utilization | > 80% |
| Pod Memory Usage | Container memory utilization | > 90% |
| Database Connections | Active DB connections | > 90% of pool |
| Redis Hit Rate | Cache hit ratio | < 80% |

### Custom Dashboards

Grafana dashboards are automatically provisioned:

- API Dashboard - API performance metrics
- Infrastructure Dashboard - Kubernetes & AWS resources
- Database Dashboard - PostgreSQL metrics
- Business Metrics - Transaction volumes, user growth

---

## 🔒 Security

### Security Best Practices

✅ Implemented Security Measures:

- All data encrypted at rest (AWS KMS)
- TLS 1.2+ for data in transit
- Regular security scanning (Trivy, Snyk)
- RBAC enabled on Kubernetes
- Network policies for pod isolation
- Secrets stored in AWS Secrets Manager
- WAF rules for DDoS protection
- Rate limiting on APIs
- Input validation and sanitization
- SQL injection prevention
- XSS protection headers
- CSRF tokens for state-changing operations

### Compliance

- PCI-DSS: Ready for Level 1 compliance
- GDPR: Data protection controls in place
- SOC 2: Audit-ready logging and monitoring

### Security Scanning

```bash
# Scan Docker images
trivy image globepay-backend:latest

# Scan dependencies
snyk test

# Run security audit
npm audit
go list -json -m all | nancy sleuth
```

---

## 📚 API Documentation

API documentation is available at `/swagger` endpoint when running the backend.

```bash
# Access Swagger UI
open http://localhost:8080/swagger/index.html
```

---

## 🤝 Contributing

We welcome contributions! Please see `CONTRIBUTING.md` for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

---

## 🙏 Acknowledgments

- React Team for the amazing frontend framework
- Go Team for the powerful backend language
- CNCF for Kubernetes and cloud-native tools
- AWS for the robust cloud infrastructure

---

## 📞 Support

- Documentation: `docs.globepay.com`
- Issues: [GitHub Issues](https://github.com/your-org/globepay/issues)
- Email: `support@globepay.com`
- Slack: Join our Slack

<div align="center">
Made with ❤️ by the Globepay Team
<br/>
<a href="https://globepay.com">Website</a> • 
<a href="https://docs.globepay.com">Documentation</a> • 
<a href="https://blog.globepay.com">Blog</a>
</div># Trigger new deployment

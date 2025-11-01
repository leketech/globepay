# Project Structure Documentation

This document provides an overview of the Globepay project structure and organization.

## Overview

The Globepay project follows a monorepo structure with clearly defined directories for different components of the application. The structure is designed to separate concerns while maintaining ease of development and deployment.

## Root Directory Structure

```
globepay/
├── backend/                 # Go backend application
├── frontend/                # React frontend application
├── infrastructure/          # Terraform infrastructure as code
├── k8s/                     # Kubernetes manifests
├── monitoring/              # Monitoring and observability configurations
├── scripts/                 # Utility scripts
├── docs/                    # Documentation files
├── .github/                 # GitHub Actions workflows
├── docker-compose.yml       # Docker Compose configuration
├── Makefile                 # Build and development commands
├── README.md                # Project overview and getting started guide
├── LICENSE                  # License information
├── CHANGELOG.md             # Version history
├── .gitignore               # Git ignore patterns
└── .editorconfig            # Editor configuration
```

## Backend Structure

The backend is a Go application following clean architecture principles.

```
backend/
├── cmd/                     # Application entry points
│   ├── api/                 # Main API server
│   │   └── main.go          # API server entry point
│   └── migration/           # Database migration tool
│       └── main.go          # Migration tool entry point
├── internal/                # Private application code
│   ├── api/                 # HTTP handlers and routes
│   │   ├── handler/         # Request handlers
│   │   ├── middleware/      # Middleware functions
│   │   └── router/          # Route definitions
│   ├── domain/              # Business logic and entities
│   │   ├── model/           # Domain models
│   │   └── service/         # Business services
│   ├── repository/          # Data access layer
│   ├── infrastructure/      # External dependencies
│   │   ├── config/          # Configuration management
│   │   ├── database/        # Database connections and migrations
│   │   ├── cache/           # Cache implementations
│   │   ├── queue/           # Message queue implementations
│   │   ├── email/           # Email service implementations
│   │   ├── sms/             # SMS service implementations
│   │   └── logger/          # Logging implementations
│   └── utils/               # Utility functions
├── test/                    # Test files
│   ├── integration/         # Integration tests
│   └── fixtures/            # Test data fixtures
├── docs/                    # Backend-specific documentation
├── .air.toml                # Air configuration for hot reload
├── Dockerfile               # Production Dockerfile
├── Dockerfile.dev           # Development Dockerfile
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── Makefile                 # Backend-specific Make commands
```

### Backend API Structure

The API layer is organized by feature:

```
backend/internal/api/
├── handler/                 # HTTP request handlers
│   ├── auth_handler.go      # Authentication endpoints
│   ├── user_handler.go      # User management endpoints
│   ├── account_handler.go   # Account management endpoints
│   ├── transfer_handler.go  # Money transfer endpoints
│   ├── transaction_handler.go # Transaction endpoints
│   └── beneficiary_handler.go # Beneficiary endpoints
├── middleware/              # Middleware functions
│   ├── auth_middleware.go   # Authentication middleware
│   ├── cors_middleware.go   # CORS middleware
│   ├── logging_middleware.go # Logging middleware
│   ├── metrics_middleware.go # Metrics middleware
│   └── rate_limit_middleware.go # Rate limiting middleware
└── router/                  # Route definitions
    └── router.go            # Main router configuration
```

### Backend Domain Structure

The domain layer contains business logic organized by entity:

```
backend/internal/domain/
├── model/                   # Domain models
│   ├── user.go              # User entity
│   ├── account.go           # Account entity
│   ├── transfer.go          # Transfer entity
│   ├── transaction.go       # Transaction entity
│   ├── beneficiary.go       # Beneficiary entity
│   └── currency.go          # Currency entity
└── service/                 # Business services
    ├── auth_service.go      # Authentication service
    ├── user_service.go      # User management service
    ├── account_service.go   # Account management service
    ├── transfer_service.go  # Money transfer service
    ├── transaction_service.go # Transaction service
    └── beneficiary_service.go # Beneficiary service
```

### Backend Repository Structure

The repository layer handles data access:

```
backend/internal/repository/
├── user_repository.go       # User data access
├── account_repository.go    # Account data access
├── transfer_repository.go   # Transfer data access
├── transaction_repository.go # Transaction data access
└── beneficiary_repository.go # Beneficiary data access
```

### Backend Infrastructure Structure

The infrastructure layer contains external service implementations:

```
backend/internal/infrastructure/
├── config/                  # Configuration management
│   └── config.go            # Configuration loader
├── database/                # Database connections and migrations
│   ├── connection.go        # Database connection
│   └── migrations/          # Database migration files
├── cache/                   # Cache implementations
│   └── redis.go             # Redis cache implementation
├── queue/                   # Message queue implementations
│   └── sqs.go               # AWS SQS implementation
├── email/                   # Email service implementations
│   └── ses.go               # AWS SES implementation
├── sms/                     # SMS service implementations
│   └── sns.go               # AWS SNS implementation
└── logger/                  # Logging implementations
    └── logger.go            # Logger configuration
```

## Frontend Structure

The frontend is a React application with TypeScript following a component-based architecture.

```
frontend/
├── public/                  # Static assets
│   ├── index.html           # HTML template
│   ├── favicon.ico          # Favicon
│   └── manifest.json        # Web app manifest
├── src/                     # Source code
│   ├── components/          # React components
│   │   ├── auth/            # Authentication components
│   │   ├── common/          # Common/shared components
│   │   ├── dashboard/       # Dashboard components
│   │   ├── transfer/        # Transfer components
│   │   └── transaction/     # Transaction components
│   ├── pages/               # Page components
│   │   ├── Auth/            # Authentication pages
│   │   ├── Dashboard/       # Dashboard pages
│   │   ├── Transfer/        # Transfer pages
│   │   └── Transaction/     # Transaction pages
│   ├── services/            # API service clients
│   ├── store/               # Redux store
│   │   ├── index.ts         # Store configuration
│   │   └── slices/          # Redux slices
│   ├── hooks/               # Custom React hooks
│   ├── utils/               # Utility functions
│   ├── types/               # TypeScript types
│   ├── constants/           # Application constants
│   ├── assets/              # Application assets
│   └── App.tsx              # Main application component
├── tests/                   # Test files
│   ├── unit/                # Unit tests
│   └── e2e/                 # End-to-end tests
├── docs/                    # Frontend-specific documentation
├── Dockerfile               # Docker configuration
├── nginx.conf               # Nginx configuration
├── package.json             # Node.js dependencies
├── tsconfig.json            # TypeScript configuration
├── vite.config.ts           # Vite configuration
└── tailwind.config.js       # Tailwind CSS configuration
```

### Frontend Components Structure

React components are organized by feature and type:

```
frontend/src/components/
├── auth/                    # Authentication components
│   ├── Login.tsx            # Login form
│   ├── Signup.tsx           # Signup form
│   └── ForgotPassword.tsx   # Forgot password form
├── common/                  # Common/shared components
│   ├── Header.tsx           # Application header
│   ├── Footer.tsx           # Application footer
│   ├── Button.tsx           # Reusable button component
│   ├── Input.tsx            # Reusable input component
│   └── Modal.tsx            # Reusable modal component
├── dashboard/               # Dashboard components
│   ├── BalanceCard.tsx      # Account balance display
│   └── RecentTransactions.tsx # Recent transactions list
├── transfer/                # Transfer components
│   ├── TransferForm.tsx     # Transfer form
│   └── BeneficiaryList.tsx  # Beneficiary selection
└── transaction/             # Transaction components
    └── TransactionList.tsx  # Transaction history list
```

### Frontend Pages Structure

Page components represent complete screens:

```
frontend/src/pages/
├── Auth/                    # Authentication pages
│   ├── LoginPage.tsx        # Login page
│   ├── SignupPage.tsx       # Signup page
│   └── ForgotPasswordPage.tsx # Forgot password page
├── Dashboard/               # Dashboard pages
│   ├── DashboardPage.tsx    # Main dashboard
│   └── ProfilePage.tsx      # User profile page
├── Transfer/                # Transfer pages
│   ├── TransferPage.tsx     # Money transfer page
│   └── BeneficiariesPage.tsx # Beneficiary management page
└── Transaction/             # Transaction pages
    └── TransactionHistoryPage.tsx # Transaction history page
```

### Frontend Services Structure

API service clients handle communication with the backend:

```
frontend/src/services/
├── api.ts                   # Base API client
├── auth.service.ts          # Authentication API service
├── user.service.ts          # User API service
├── account.service.ts       # Account API service
├── transfer.service.ts      # Transfer API service
└── transaction.service.ts   # Transaction API service
```

### Frontend Store Structure

Redux store manages application state:

```
frontend/src/store/
├── index.ts                 # Store configuration
└── slices/                  # Redux slices
    ├── authSlice.ts         # Authentication state
    ├── userSlice.ts         # User state
    ├── accountSlice.ts      # Account state
    ├── transferSlice.ts     # Transfer state
    └── transactionSlice.ts  # Transaction state
```

## Infrastructure Structure

Terraform configurations for infrastructure as code:

```
infrastructure/
├── modules/                 # Reusable Terraform modules
│   ├── networking/          # VPC and networking module
│   ├── eks/                 # EKS cluster module
│   ├── rds/                 # RDS database module
│   ├── elasticache/         # ElastiCache module
│   ├── s3/                  # S3 bucket module
│   ├── cloudfront/          # CloudFront CDN module
│   ├── sqs/                 # SQS queue module
│   ├── ses/                 # SES email module
│   └── kms/                 # KMS key module
├── environments/            # Environment-specific configurations
│   ├── dev/                 # Development environment
│   │   ├── main.tf          # Main configuration
│   │   ├── variables.tf     # Variable definitions
│   │   ├── terraform.tfvars # Variable values
│   │   └── outputs.tf       # Output definitions
│   ├── staging/             # Staging environment
│   │   ├── main.tf          # Main configuration
│   │   ├── variables.tf     # Variable definitions
│   │   ├── terraform.tfvars # Variable values
│   │   └── outputs.tf       # Output definitions
│   └── prod/                # Production environment
│       ├── main.tf          # Main configuration
│       ├── variables.tf     # Variable definitions
│       ├── terraform.tfvars # Variable values
│       └── outputs.tf       # Output definitions
├── docs/                    # Infrastructure documentation
└── README.md                # Infrastructure overview
```

## Kubernetes Structure

Kubernetes manifests for application deployment:

```
k8s/
├── base/                    # Base Kubernetes configurations
│   ├── backend/             # Backend deployment
│   │   ├── deployment.yaml  # Backend deployment
│   │   ├── service.yaml     # Backend service
│   │   └── hpa.yaml         # Backend horizontal pod autoscaler
│   ├── frontend/            # Frontend deployment
│   │   ├── deployment.yaml  # Frontend deployment
│   │   ├── service.yaml     # Frontend service
│   │   └── ingress.yaml     # Frontend ingress
│   ├── database/            # Database deployment
│   │   └── service.yaml     # Database service
│   └── kustomization.yaml   # Base kustomization
├── overlays/                # Environment overlays
│   ├── dev/                 # Development overlay
│   │   ├── kustomization.yaml # Dev kustomization
│   │   └── patches/         # Dev-specific patches
│   ├── staging/             # Staging overlay
│   │   ├── kustomization.yaml # Staging kustomization
│   │   └── patches/         # Staging-specific patches
│   └── prod/                # Production overlay
│       ├── kustomization.yaml # Prod kustomization
│       └── patches/         # Prod-specific patches
├── argocd/                  # ArgoCD configurations
│   ├── application.yaml     # Application definition
│   └── project.yaml         # Project definition
└── docs/                    # Kubernetes documentation
```

## Monitoring Structure

Monitoring and observability configurations:

```
monitoring/
├── prometheus/              # Prometheus configurations
│   ├── prometheus-config.yaml # Prometheus configuration
│   └── rules/               # Alert rules
├── grafana/                 # Grafana configurations
│   ├── dashboards/          # Dashboard definitions
│   └── datasources/         # Data source configurations
├── loki/                    # Loki configurations
│   └── loki-config.yaml     # Loki configuration
├── promtail/                # Promtail configurations
│   └── promtail-config.yaml # Promtail configuration
├── jaeger/                  # Jaeger configurations
│   └── jaeger-config.yaml   # Jaeger configuration
├── alertmanager/            # Alertmanager configurations
│   └── alertmanager-config.yaml # Alertmanager configuration
└── docs/                    # Monitoring documentation
```

## Scripts Structure

Utility scripts for development and deployment:

```
scripts/
├── setup-dev-environment.sh # Development environment setup
├── deploy.sh                # Deployment script
├── backup.sh                # Backup script
├── restore.sh               # Restore script
├── migrate.sh               # Database migration script
└── docs/                    # Script documentation
```

## Documentation Structure

Project documentation files:

```
docs/
├── quickstart.md            # Quick start guide
├── architecture.md          # Architecture documentation
├── api.md                   # API documentation
├── api-reference.md         # Detailed API reference
├── database.md              # Database schema documentation
├── testing.md               # Testing documentation
├── monitoring.md            # Monitoring documentation
├── security.md              # Security documentation
├── deployment.md            # Deployment documentation
├── troubleshooting.md       # Troubleshooting guide
├── glossary.md              # Glossary of terms
├── project-structure.md     # Project structure documentation
└── CHANGELOG.md             # Version history
```

## GitHub Actions Structure

CI/CD workflows:

```
.github/
├── workflows/               # Workflow definitions
│   ├── test.yml             # Test workflow
│   ├── build.yml            # Build workflow
│   ├── deploy-dev.yml       # Development deployment workflow
│   ├── deploy-staging.yml   # Staging deployment workflow
│   ├── deploy-prod.yml      # Production deployment workflow
│   └── security-scan.yml    # Security scanning workflow
└── docs/                    # GitHub Actions documentation
```

## Configuration Files

Important configuration files at the root level:

- **docker-compose.yml**: Docker Compose configuration for local development
- **Makefile**: Build and development commands
- **.gitignore**: Git ignore patterns
- **.editorconfig**: Editor configuration for consistent coding styles
- **.air.toml**: Air configuration for Go hot reloading
- **.dockerignore**: Docker ignore patterns
- **.env.example**: Example environment variables

This project structure documentation provides a comprehensive overview of how the Globepay application is organized, making it easier for developers to navigate and contribute to the codebase.
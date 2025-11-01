# Developer Quick Start Guide

This guide provides step-by-step instructions for developers to get started with the Globepay project.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Git** - Version control system
- **Docker** (24.0+) - Container runtime
- **Docker Compose** (2.20+) - Multi-container Docker applications
- **Go** (1.21+) - Backend programming language
- **Node.js** (20+) - Frontend runtime environment
- **npm** (9+) - Node package manager
- **kubectl** (1.28+) - Kubernetes command-line tool
- **AWS CLI** (2.13+) - AWS command-line interface
- **Terraform** (1.5+) - Infrastructure as code tool

## Getting the Code

1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/globepay.git
   cd globepay
   ```

2. Check out the development branch:
   ```bash
   git checkout develop
   ```

## Setting Up the Development Environment

### Automated Setup

Run the setup script to automatically configure your development environment:

```bash
# Make the script executable
chmod +x scripts/setup-dev-environment.sh

# Run the setup script
./scripts/setup-dev-environment.sh
```

### Manual Setup

If you prefer to set up manually:

1. **Start Docker services:**
   ```bash
   docker-compose up -d
   ```

2. **Set up the backend:**
   ```bash
   cd backend
   go mod download
   cd ..
   ```

3. **Set up the frontend:**
   ```bash
   cd frontend
   npm install
   cd ..
   ```

4. **Run database migrations:**
   ```bash
   make migration-up
   ```

## Running the Application

### Using Docker Compose (Recommended)

Start all services with Docker Compose:

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Running Services Individually

#### Backend

```bash
# Navigate to backend directory
cd backend

# Run with hot reload (requires air)
air

# Or run directly
go run cmd/api/main.go

# Run tests
go test ./...
```

#### Frontend

```bash
# Navigate to frontend directory
cd frontend

# Start development server
npm run dev

# Run tests
npm test

# Build for production
npm run build
```

## Accessing Services

Once the application is running, you can access the following services:

| Service | URL | Description |
|---------|-----|-------------|
| Frontend | http://localhost:3000 | React application |
| Backend API | http://localhost:8080 | Go REST API |
| Swagger Docs | http://localhost:8080/swagger | API documentation |
| PostgreSQL | localhost:5432 | Database |
| Redis | localhost:6379 | Cache |
| Grafana | http://localhost:3001 | Monitoring dashboard |
| Prometheus | http://localhost:9091 | Metrics collection |

## Development Workflow

### Backend Development

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make changes to the code:**
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation if needed

3. **Run tests:**
   ```bash
   cd backend
   go test ./...
   ```

4. **Commit changes:**
   ```bash
   git add .
   git commit -m "Add feature: your feature description"
   ```

### Frontend Development

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make changes to the code:**
   - Follow TypeScript and React best practices
   - Add tests for new components
   - Update documentation if needed

3. **Run tests:**
   ```bash
   cd frontend
   npm test
   ```

4. **Commit changes:**
   ```bash
   git add .
   git commit -m "Add feature: your feature description"
   ```

## Database Management

### Running Migrations

```bash
# Run all pending migrations
make migration-up

# Rollback last migration
make migration-down

# Create new migration
make migration-create NAME=your_migration_name
```

### Accessing the Database

```bash
# Connect to PostgreSQL
docker-compose exec postgres psql -U postgres -d globepay

# View database schema
\d
```

## Testing

### Backend Testing

```bash
# Run all backend tests
make test-backend

# Run specific test
cd backend
go test ./internal/service -run TestTransferService

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Frontend Testing

```bash
# Run all frontend tests
make test-frontend

# Run specific test
cd frontend
npm test src/components/auth/Login.test.tsx

# Run with coverage
npm run test:coverage
```

### End-to-End Testing

```bash
# Run E2E tests
cd frontend
npm run test:e2e

# Run E2E tests in UI mode
npm run test:e2e:ui
```

## Code Quality

### Linting

```bash
# Backend linting
cd backend
golangci-lint run

# Frontend linting
cd frontend
npm run lint
```

### Formatting

```bash
# Backend formatting
cd backend
go fmt ./...

# Frontend formatting
cd frontend
npm run format
```

## Monitoring and Debugging

### Accessing Monitoring Tools

```bash
# Port-forward Grafana
kubectl port-forward -n monitoring svc/prometheus-grafana 3000:80

# Port-forward Prometheus
kubectl port-forward -n monitoring svc/prometheus-kube-prometheus-prometheus 9090:9090
```

### Viewing Logs

```bash
# View all service logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f backend

# View Kubernetes logs
kubectl logs -n globepay-prod -l app=backend
```

## Common Development Tasks

### Adding a New API Endpoint

1. Create a new handler in `backend/internal/api/handler/`
2. Add the route in `backend/internal/api/router/router.go`
3. Create a service method in `backend/internal/domain/service/`
4. Create a repository method in `backend/internal/repository/`
5. Add tests for the new functionality

### Adding a New Frontend Component

1. Create a new component in `frontend/src/components/`
2. Add tests in `frontend/src/components/__tests__/`
3. Import and use the component in a page
4. Add stories for Storybook (if applicable)

### Adding a New Database Migration

```bash
# Create migration files
make migration-create NAME=add_users_table

# Edit the generated .up.sql and .down.sql files
# Run the migration
make migration-up
```

## Environment Variables

### Backend

Create `backend/.env`:
```env
ENVIRONMENT=development
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/globepay?sslmode=disable
REDIS_URL=redis://localhost:6379/0
JWT_SECRET=your-jwt-secret
AWS_REGION=us-east-1
```

### Frontend

Create `frontend/.env`:
```env
VITE_API_URL=http://localhost:8080
VITE_ENVIRONMENT=development
```

## Troubleshooting

### Common Issues

1. **Services not starting:**
   - Check Docker daemon is running
   - Verify port availability
   - Check Docker Compose logs

2. **Database connection issues:**
   - Verify database is running
   - Check connection string
   - Ensure migrations are run

3. **Frontend not loading:**
   - Check if backend is running
   - Verify API URL in environment variables
   - Check browser console for errors

### Useful Commands

```bash
# Check service status
docker-compose ps

# Restart a specific service
docker-compose restart backend

# Rebuild services
docker-compose up --build

# Clean up Docker resources
docker system prune -a
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Update documentation
6. Submit a pull request

For more detailed information, see the [CONTRIBUTING.md](../CONTRIBUTING.md) file.

## Getting Help

If you need help:

1. Check the documentation in the `docs/` directory
2. Review existing issues on GitHub
3. Ask questions in the development Slack channel
4. Contact the maintainers

This quick start guide should get you up and running with the Globepay development environment. For more detailed information about specific components, refer to the individual documentation files in the `docs/` directory.
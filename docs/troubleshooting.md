# Troubleshooting Guide

This document provides solutions to common issues encountered when working with the Globepay application.

## Table of Contents

- [Development Environment Issues](#development-environment-issues)
- [Docker and Container Issues](#docker-and-container-issues)
- [Database Issues](#database-issues)
- [API Issues](#api-issues)
- [Frontend Issues](#frontend-issues)
- [Kubernetes Issues](#kubernetes-issues)
- [Monitoring Issues](#monitoring-issues)
- [Deployment Issues](#deployment-issues)
- [Security Issues](#security-issues)
- [Performance Issues](#performance-issues)

## Development Environment Issues

### Go Installation Problems

**Issue**: Go commands not found or version mismatch

**Solution**:
```bash
# Check Go version
go version

# If Go is not installed or wrong version:
# Download and install Go 1.21+ from https://golang.org/dl/
# Or use package manager:
# Ubuntu/Debian: sudo apt install golang
# macOS: brew install go
# Windows: choco install golang

# Verify installation
go version
```

### Node.js Installation Problems

**Issue**: Node.js commands not found or version mismatch

**Solution**:
```bash
# Check Node.js version
node --version
npm --version

# If Node.js is not installed or wrong version:
# Download and install Node.js 20+ from https://nodejs.org/
# Or use package manager:
# Ubuntu/Debian: sudo apt install nodejs npm
# macOS: brew install node
# Windows: choco install nodejs

# Verify installation
node --version
npm --version
```

### Environment Variables Not Loading

**Issue**: Environment variables not being read by the application

**Solution**:
```bash
# Check if .env file exists
ls -la backend/.env
ls -la frontend/.env

# Verify environment variable loading
# Backend:
cd backend
go run cmd/api/main.go
# Check logs for "Environment variables loaded successfully"

# Frontend:
cd frontend
echo $VITE_API_URL
# Should show the API URL

# If variables are not loading, check:
# 1. .env file format (KEY=VALUE)
# 2. No spaces around equals sign
# 3. Proper quoting for values with spaces
# 4. File permissions
```

## Docker and Container Issues

### Docker Daemon Not Running

**Issue**: Docker commands fail with "Cannot connect to the Docker daemon"

**Solution**:
```bash
# Linux:
sudo systemctl start docker
sudo systemctl enable docker

# macOS:
# Start Docker Desktop application

# Windows:
# Start Docker Desktop application

# Verify Docker is running:
docker version
docker info
```

### Docker Compose Services Not Starting

**Issue**: Services fail to start or exit immediately

**Solution**:
```bash
# Check service status
docker-compose ps

# View service logs
docker-compose logs <service-name>
docker-compose logs backend
docker-compose logs frontend

# Check for port conflicts
lsof -i :8080
lsof -i :3000

# Rebuild services
docker-compose down
docker-compose up --build

# Check container resource limits
docker stats
```

### Container Build Failures

**Issue**: Docker build process fails

**Solution**:
```bash
# Check Dockerfile syntax
docker run --rm -v $(pwd):/project docker/dockerfile:latest /project/Dockerfile

# Build with verbose output
docker build --no-cache -t globepay-backend .

# Check build context size
du -sh .

# Clean up Docker resources
docker system prune -a
docker volume prune
```

### Volume Mounting Issues

**Issue**: Files not accessible in containers or permissions errors

**Solution**:
```bash
# Check volume mounts
docker-compose config

# Verify file permissions
ls -la backend/
ls -la frontend/

# Fix permissions (Linux/macOS)
sudo chown -R $(id -u):$(id -g) backend/
sudo chown -R $(id -u):$(id -g) frontend/

# Check Docker Desktop settings (Windows/macOS)
# Ensure file sharing is enabled for project directories
```

## Database Issues

### PostgreSQL Connection Failed

**Issue**: Application cannot connect to PostgreSQL database

**Solution**:
```bash
# Check if PostgreSQL container is running
docker-compose ps postgres

# Test database connection
docker-compose exec postgres pg_isready

# Check database logs
docker-compose logs postgres

# Verify connection string
# Format: postgresql://username:password@host:port/database
echo "postgresql://postgres:postgres@localhost:5432/globepay"

# Test connection manually
psql -h localhost -p 5432 -U postgres -d globepay

# Check if database exists
docker-compose exec postgres psql -U postgres -l
```

### Database Migration Failures

**Issue**: Database migrations fail to run

**Solution**:
```bash
# Check migration status
docker-compose exec backend /app/migrate version

# Run migrations manually
docker-compose exec backend /app/migrate up

# Check migration files
ls -la backend/internal/infrastructure/database/migrations/

# Reset migrations (development only)
docker-compose exec backend /app/migrate down -all
docker-compose exec backend /app/migrate up

# Check database schema
docker-compose exec postgres psql -U postgres -d globepay -c "\dt"
```

### Slow Database Queries

**Issue**: Database queries are taking too long

**Solution**:
```bash
# Enable query logging
docker-compose exec postgres psql -U postgres -d globepay -c "ALTER SYSTEM SET log_min_duration_statement = 1000;"
docker-compose exec postgres psql -U postgres -d globepay -c "SELECT pg_reload_conf();"

# Check slow query logs
docker-compose exec postgres cat /var/lib/postgresql/data/log/postgresql-*.log

# Analyze query performance
docker-compose exec postgres psql -U postgres -d globepay -c "EXPLAIN ANALYZE SELECT * FROM users WHERE email = 'test@example.com';"

# Add indexes for frequently queried columns
docker-compose exec postgres psql -U postgres -d globepay -c "CREATE INDEX CONCURRENTLY idx_users_email ON users(email);"
```

### Database Connection Pool Exhaustion

**Issue**: "Too many connections" errors

**Solution**:
```bash
# Check current connections
docker-compose exec postgres psql -U postgres -d globepay -c "SELECT count(*) FROM pg_stat_activity;"

# Check connection limits
docker-compose exec postgres psql -U postgres -d globepay -c "SHOW max_connections;"

# Terminate idle connections
docker-compose exec postgres psql -U postgres -d globepay -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE state = 'idle' AND pid <> pg_backend_pid();"

# Increase connection pool size in application config
# backend/.env:
# DB_MAX_OPEN_CONNS=50
# DB_MAX_IDLE_CONNS=10
```

## API Issues

### API Endpoints Returning 404

**Issue**: API endpoints return 404 Not Found

**Solution**:
```bash
# Check if backend is running
curl -v http://localhost:8080/health

# Check API routes
curl http://localhost:8080/api/v1/health

# Check backend logs
docker-compose logs backend

# Verify router configuration
# Check backend/internal/api/router/router.go

# Restart backend service
docker-compose restart backend
```

### Authentication Failures

**Issue**: Login returns 401 Unauthorized

**Solution**:
```bash
# Check user exists in database
docker-compose exec postgres psql -U postgres -d globepay -c "SELECT email, password_hash FROM users WHERE email = 'test@example.com';"

# Verify password hashing
# Passwords should be bcrypt hashed

# Check JWT configuration
# backend/.env:
# JWT_SECRET=your-secret-key
# JWT_EXPIRATION=24h

# Test token generation
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Check token validity
# Decode JWT at https://jwt.io/
```

### CORS Errors

**Issue**: Frontend cannot make API requests due to CORS

**Solution**:
```bash
# Check CORS middleware configuration
# backend/internal/api/middleware/cors.go

# Verify allowed origins
# Should include http://localhost:3000 for development

# Check browser console for CORS errors
# Open Developer Tools -> Console

# Test CORS headers
curl -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: X-Requested-With" \
  -X OPTIONS \
  http://localhost:8080/api/v1/auth/login

# Add CORS headers if missing
# In backend/internal/api/middleware/cors.go:
# c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
# c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
# c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
```

### Rate Limiting Issues

**Issue**: Requests are being rate limited

**Solution**:
```bash
# Check rate limiting configuration
# backend/internal/api/middleware/ratelimit.go

# Verify rate limit settings
# Default: 100 requests per minute

# Check rate limit headers in response
curl -v http://localhost:8080/api/v1/health

# Look for:
# X-RateLimit-Limit
# X-RateLimit-Remaining
# X-RateLimit-Reset

# Temporarily disable rate limiting for testing
# Comment out rate limit middleware in router
```

## Frontend Issues

### Frontend Not Loading

**Issue**: Frontend application fails to start or shows blank page

**Solution**:
```bash
# Check frontend container status
docker-compose ps frontend

# Check frontend logs
docker-compose logs frontend

# Verify build process
docker-compose exec frontend ls -la /usr/share/nginx/html

# Check nginx configuration
docker-compose exec frontend cat /etc/nginx/nginx.conf

# Test static file serving
curl http://localhost:3000/index.html

# Rebuild frontend
docker-compose build frontend
docker-compose up -d frontend
```

### React Component Errors

**Issue**: React components fail to render or show errors

**Solution**:
```bash
# Check browser console for errors
# Open Developer Tools -> Console

# Check React component structure
# Verify props and state management

# Check Redux store
# Verify actions and reducers

# Check API service calls
# Verify API endpoints and response handling

# Enable React development tools
# Install React Developer Tools browser extension
```

### Build Failures

**Issue**: Frontend build process fails

**Solution**:
```bash
# Check build logs
docker-compose logs frontend

# Check package.json dependencies
docker-compose exec frontend cat package.json

# Install missing dependencies
docker-compose exec frontend npm install

# Check TypeScript compilation errors
docker-compose exec frontend npm run build

# Clear build cache
docker-compose exec frontend rm -rf node_modules/.cache
docker-compose exec frontend npm run build
```

### Environment Variables Not Loading

**Issue**: VITE environment variables not available in frontend

**Solution**:
```bash
# Check .env file format
# Variables must be prefixed with VITE_
# frontend/.env:
# VITE_API_URL=http://localhost:8080
# VITE_ENVIRONMENT=development

# Restart frontend service
docker-compose restart frontend

# Check if variables are injected
docker-compose exec frontend printenv | grep VITE_

# Verify usage in code
# Use import.meta.env.VITE_API_URL in components
```

## Kubernetes Issues

### Pod CrashLoopBackOff

**Issue**: Pods are stuck in CrashLoopBackOff status

**Solution**:
```bash
# Check pod status
kubectl get pods -n globepay-prod

# Check pod logs
kubectl logs -n globepay-prod <pod-name>

# Check pod description
kubectl describe pod -n globepay-prod <pod-name>

# Check events
kubectl get events -n globepay-prod

# Check resource limits
kubectl describe limitrange -n globepay-prod

# Check if images exist
kubectl describe deployment -n globepay-prod backend
```

### Service Not Accessible

**Issue**: Kubernetes services are not reachable

**Solution**:
```bash
# Check service status
kubectl get services -n globepay-prod

# Check service endpoints
kubectl get endpoints -n globepay-prod

# Check if pods are ready
kubectl get pods -n globepay-prod -o wide

# Test service connectivity
kubectl port-forward -n globepay-prod svc/backend 8080:80
curl http://localhost:8080/health

# Check network policies
kubectl get networkpolicies -n globepay-prod
```

### ConfigMap/Secret Issues

**Issue**: Configuration or secrets not loaded properly

**Solution**:
```bash
# Check ConfigMaps
kubectl get configmaps -n globepay-prod
kubectl describe configmap -n globepay-prod backend-config

# Check Secrets
kubectl get secrets -n globepay-prod
kubectl describe secret -n globepay-prod backend-secrets

# Verify volume mounts
kubectl describe pod -n globepay-prod <pod-name>

# Check environment variables
kubectl exec -it -n globepay-prod <pod-name> -- printenv
```

### Horizontal Pod Autoscaler Not Working

**Issue**: HPA is not scaling pods

**Solution**:
```bash
# Check HPA status
kubectl get hpa -n globepay-prod

# Check HPA details
kubectl describe hpa -n globepay-prod backend-hpa

# Check metrics server
kubectl top nodes
kubectl top pods -n globepay-prod

# Check if metrics are available
kubectl get --raw "/apis/metrics.k8s.io/v1beta1/namespaces/globepay-prod/pods"

# Verify resource requests in deployment
kubectl describe deployment -n globepay-prod backend
```

## Monitoring Issues

### Prometheus Not Scraping Metrics

**Issue**: Prometheus is not collecting metrics from targets

**Solution**:
```bash
# Check Prometheus targets
kubectl port-forward -n monitoring svc/prometheus-kube-prometheus-prometheus 9090:9090
# Visit http://localhost:9090/targets

# Check service discovery
# Visit http://localhost:9090/service-discovery

# Check if metrics endpoint is accessible
curl http://backend.globepay-prod.svc.cluster.local:9090/metrics

# Check Prometheus configuration
kubectl get configmaps -n monitoring prometheus-kube-prometheus-prometheus
```

### Grafana Dashboards Not Loading

**Issue**: Grafana dashboards show no data or fail to load

**Solution**:
```bash
# Check Grafana status
kubectl get pods -n monitoring -l app.kubernetes.io/name=grafana

# Check Grafana logs
kubectl logs -n monitoring -l app.kubernetes.io/name=grafana

# Verify data source configuration
# Visit http://localhost:3000 (admin/prom-operator)
# Configuration -> Data Sources -> Prometheus

# Check dashboard provisioning
kubectl get configmaps -n monitoring | grep dashboard

# Restart Grafana
kubectl delete pod -n monitoring -l app.kubernetes.io/name=grafana
```

### Alertmanager Not Sending Notifications

**Issue**: Alerts are not being sent to notification channels

**Solution**:
```bash
# Check Alertmanager status
kubectl get pods -n monitoring -l app=alertmanager

# Check Alertmanager configuration
kubectl get configmaps -n monitoring alertmanager-alertmanager

# Check alert rules
kubectl get prometheusrules -n monitoring

# Test alert firing
# Visit http://localhost:9090/alerts

# Check Alertmanager UI
kubectl port-forward -n monitoring svc/alertmanager-operated 9093:9093
# Visit http://localhost:9093
```

### Loki Log Collection Issues

**Issue**: Logs are not appearing in Loki

**Solution**:
```bash
# Check Promtail pods
kubectl get pods -n monitoring -l app=promtail

# Check Promtail logs
kubectl logs -n monitoring -l app=promtail

# Check Promtail configuration
kubectl get configmaps -n monitoring promtail-promtail

# Verify log file paths
kubectl exec -it -n monitoring -l app=promtail -- cat /etc/promtail/promtail.yaml

# Test Loki connectivity
kubectl port-forward -n monitoring svc/loki 3100:3100
curl http://localhost:3100/ready
```

## Deployment Issues

### Terraform Apply Failures

**Issue**: Terraform fails to apply infrastructure changes

**Solution**:
```bash
# Check Terraform state
cd infrastructure/environments/prod
terraform state list

# Check for errors in plan
terraform plan

# Verify AWS credentials
aws sts get-caller-identity

# Check Terraform version
terraform version

# Initialize backend
terraform init -reconfigure

# Apply with detailed output
terraform apply -auto-approve
```

### ArgoCD Sync Failures

**Issue**: ArgoCD applications are out of sync or failing

**Solution**:
```bash
# Check ArgoCD application status
kubectl get applications -n argocd

# Check application details
kubectl describe application -n argocd globepay-backend-prod

# Check ArgoCD logs
kubectl logs -n argocd -l app.kubernetes.io/name=argocd-server

# Force sync application
argocd app sync globepay-backend-prod

# Check for health issues
argocd app health globepay-backend-prod
```

### Docker Image Push Failures

**Issue**: Cannot push Docker images to container registry

**Solution**:
```bash
# Check ECR login
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 907849381252.dkr.ecr.us-east-1.amazonaws.com

# Verify repository exists
aws ecr describe-repositories --repository-names globepay-backend

# Check image tags
docker images globepay-backend

# Push with verbose output
docker push 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend:latest --verbose
```

### Helm Chart Installation Failures

**Issue**: Helm charts fail to install or upgrade

**Solution**:
```bash
# Check Helm repositories
helm repo list

# Update repositories
helm repo update

# Check chart values
helm template prometheus prometheus-community/kube-prometheus-stack --values monitoring/prometheus/values.yaml

# Install with debug output
helm install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --values monitoring/prometheus/values.yaml \
  --debug

# Check release status
helm list -n monitoring
```

## Security Issues

### SSL/TLS Certificate Issues

**Issue**: SSL/TLS certificate errors or expiration

**Solution**:
```bash
# Check certificate expiration
echo | openssl s_client -connect api.globepay.com:443 2>/dev/null | openssl x509 -noout -dates

# Check certificate chain
openssl s_client -connect api.globepay.com:443 -showcerts

# Renew certificates
# For ACM certificates, they auto-renew 60 days before expiration

# Check ingress TLS configuration
kubectl describe ingress -n globepay-prod globepay-ingress
```

### JWT Token Issues

**Issue**: JWT tokens are invalid or expired

**Solution**:
```bash
# Check JWT secret configuration
kubectl get secrets -n globepay-prod backend-secrets -o yaml

# Verify token expiration settings
# Check backend configuration for JWT expiration

# Test token generation
curl -X POST https://api.globepay.com/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Decode token at https://jwt.io/
```

### Security Scanner Alerts

**Issue**: Security scanners report vulnerabilities

**Solution**:
```bash
# Run Trivy scan
trivy image globepay-backend:latest

# Run Snyk test
snyk test

# Update dependencies
cd backend && go get -u ./...
cd frontend && npm update

# Rebuild images
docker build -t globepay-backend:latest ./backend
docker build -t globepay-frontend:latest ./frontend

# Rescan to verify fixes
trivy image globepay-backend:latest
```

## Performance Issues

### High API Latency

**Issue**: API responses are slow

**Solution**:
```bash
# Check response times
curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/api/v1/health"

# Create curl-format.txt:
#     time_namelookup:  %{time_namelookup}\n
#        time_connect:  %{time_connect}\n
#     time_appconnect:  %{time_appconnect}\n
#    time_pretransfer:  %{time_pretransfer}\n
#       time_redirect:  %{time_redirect}\n
#  time_starttransfer:  %{time_starttransfer}\n
#                     ----------\n
#          time_total:  %{time_total}\n

# Check backend logs for slow queries
docker-compose logs backend | grep "slow"

# Profile Go application
go tool pprof http://localhost:8080/debug/pprof/profile

# Check database performance
docker-compose exec postgres psql -U postgres -d globepay -c "EXPLAIN ANALYZE SELECT * FROM users;"
```

### High Memory Usage

**Issue**: Application consuming excessive memory

**Solution**:
```bash
# Check container memory usage
docker stats

# Check Kubernetes resource usage
kubectl top pods -n globepay-prod

# Check Go memory profiling
go tool pprof http://localhost:8080/debug/pprof/heap

# Check for memory leaks
# Look for continuously increasing memory usage

# Adjust resource limits
# Update k8s/base/backend/deployment.yaml:
# resources:
#   requests:
#     memory: "256Mi"
#   limits:
#     memory: "512Mi"
```

### High CPU Usage

**Issue**: Application consuming excessive CPU

**Solution**:
```bash
# Check container CPU usage
docker stats

# Check Kubernetes resource usage
kubectl top nodes
kubectl top pods -n globepay-prod

# Profile CPU usage
go tool pprof http://localhost:8080/debug/pprof/profile

# Check for infinite loops or expensive operations
# Review code for optimization opportunities

# Adjust resource limits
# Update k8s/base/backend/deployment.yaml:
# resources:
#   requests:
#     cpu: "250m"
#   limits:
#     cpu: "500m"
```

### Database Performance Issues

**Issue**: Database queries are slow

**Solution**:
```bash
# Enable slow query logging
docker-compose exec postgres psql -U postgres -d globepay -c "ALTER SYSTEM SET log_min_duration_statement = 1000;"

# Check slow queries
docker-compose exec postgres cat /var/lib/postgresql/data/log/postgresql-*.log

# Analyze query performance
docker-compose exec postgres psql -U postgres -d globepay -c "EXPLAIN ANALYZE SELECT * FROM transfers WHERE user_id = 'user123';"

# Add missing indexes
docker-compose exec postgres psql -U postgres -d globepay -c "CREATE INDEX CONCURRENTLY idx_transfers_user_id ON transfers(user_id);"

# Check database statistics
docker-compose exec postgres psql -U postgres -d globepay -c "SELECT * FROM pg_stat_user_tables;"
```

## General Debugging Tips

### Enable Debug Logging

```bash
# Backend debug mode
# Set in backend/.env:
# DEBUG=true

# Frontend debug mode
# Set in frontend/.env:
# VITE_DEBUG=true

# Kubernetes debug mode
# Add --v=4 to kubectl commands
kubectl get pods -n globepay-prod --v=4
```

### Check System Resources

```bash
# Check disk space
df -h

# Check memory usage
free -h

# Check CPU usage
top

# Check network connectivity
ping google.com
```

### Useful Diagnostic Commands

```bash
# Check all running containers
docker ps -a

# Check Docker networks
docker network ls

# Check Docker volumes
docker volume ls

# Check system logs
journalctl -u docker.service

# Check Kubernetes cluster info
kubectl cluster-info

# Check Kubernetes nodes
kubectl get nodes
```

This troubleshooting guide covers the most common issues you may encounter when working with the Globepay application. If you encounter an issue not covered here, please check the logs and consider reaching out to the development team for assistance.
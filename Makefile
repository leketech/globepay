.PHONY: help setup dev-up dev-down test build deploy clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-25s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## Setup development environment
	@./scripts/setup-dev-environment.sh

dev-up: ## Start development environment
	@docker-compose up -d
	@echo "✅ Development environment is running"

dev-down: ## Stop development environment
	@docker-compose down
	@echo "✅ Development environment stopped"

dev-logs: ## Show development logs
	@docker-compose logs -f

test-backend: ## Run backend tests
	@cd backend && go test -v -race -coverprofile=coverage.out ./...

test-backend-unit: ## Run backend unit tests
	@cd backend && go test -v ./test/unit/...

test-backend-integration: ## Run backend integration tests
	@cd backend && go test -v ./test/integration/...

test-backend-coverage: ## Run backend tests with coverage
	@cd backend && go test -v -race -coverprofile=coverage.out ./test/...
	@cd backend && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at backend/coverage.html"

test-frontend: ## Run frontend tests
	@cd frontend && npm test

test: test-backend test-frontend ## Run all tests

build-backend: ## Build backend Docker image
	@docker build -t globepay-backend:latest ./backend

build-frontend: ## Build frontend Docker image
	@docker build -t globepay-frontend:latest ./frontend

build: build-backend build-frontend ## Build all Docker images

lint-backend: ## Lint backend code
	@cd backend && golangci-lint run

lint-frontend: ## Lint frontend code
	@cd frontend && npm run lint

lint: lint-backend lint-frontend ## Lint all code

deploy-staging-docker: ## Deploy to staging using Docker Compose
	@./scripts/deploy-staging.sh

deploy-prod-docker: ## Deploy to production using Docker Compose (testing only)
	@./scripts/deploy-prod.sh

deploy-dev: ## Deploy to development
	@kubectl apply -k k8s/overlays/dev

deploy-staging: ## Deploy to staging
	@kubectl apply -k k8s/overlays/staging

deploy-prod: ## Deploy to production
	@kubectl apply -k k8s/overlays/prod

deploy-prod-k8s: ## Deploy to production using Kubernetes (build, push, deploy)
	@./scripts/deploy-to-k8s.sh prod latest

deploy-staging-k8s: ## Deploy to staging using Kubernetes (build, push, deploy)
	@./scripts/deploy-to-k8s.sh staging latest

deploy-dev-k8s: ## Deploy to development using Kubernetes (build, push, deploy)
	@./scripts/deploy-to-k8s.sh dev latest

deploy-all: ## Deploy to all environments (dev -> staging -> prod)
	@./scripts/promote-changes.sh all

clean: ## Clean up development environment
	@docker-compose down -v
	@docker system prune -f
	@echo "✅ Cleaned up"

migration-create: ## Create new database migration (usage: make migration-create NAME=create_users_table)
	@cd backend && go run cmd/migration/main.go create $(NAME)

migration-up: ## Run database migrations up
	@cd backend && go run cmd/migration/main.go up

migration-down: ## Run database migrations down
	@cd backend && go run cmd/migration/main.go down

verify-setup: ## Verify project setup
	@./scripts/verify-setup.sh
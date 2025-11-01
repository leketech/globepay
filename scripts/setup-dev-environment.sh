#!/bin/bash

set -e

echo "🚀 Setting up Globepay development environment..."

# Check prerequisites
if ! command -v docker &> /dev/null
then
    echo "❌ Docker is required but not installed. Aborting."
    exit 1
fi

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p backend/internal/infrastructure/database/migrations
mkdir -p frontend/src
mkdir -p monitoring/prometheus
mkdir -p monitoring/grafana/dashboards
mkdir -p monitoring/grafana/datasources
mkdir -p k8s/base
mkdir -p k8s/overlays/{dev,staging,prod}

# Setup backend
echo "🔨 Setting up backend..."
cd backend
if [ ! -f go.mod ]; then
    /usr/local/go/bin/go mod init globepay
    /usr/local/go/bin/go mod tidy
fi
cd ..

# Setup frontend
echo "⚛️  Setting up frontend..."
cd frontend
if [ ! -f package.json ]; then
    npm init -y
    npm install react react-dom react-router-dom @reduxjs/toolkit react-redux axios
    npm install -D typescript @types/react @types/react-dom @types/node vite @vitejs/plugin-react tailwindcss postcss autoprefixer
fi
cd ..

# Create .env files if they don't exist
echo "⚙️  Creating environment files..."
if [ ! -f backend/.env ]; then
    cp backend/.env.example backend/.env
fi

if [ ! -f frontend/.env ]; then
    cp frontend/.env.example frontend/.env
fi

# Start services
echo "🐳 Starting Docker services..."
docker compose up -d

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 10

# Run database migrations
echo "📊 Running database migrations..."
cd backend
/usr/local/go/bin/go run cmd/migration/main.go up
cd ..

echo "✅ Development environment setup complete!"
echo ""
echo "📝 Services running:"
echo "  - Backend API:     http://localhost:8080"
echo "  - Frontend:        http://localhost:3000"
echo "  - Prometheus:      http://localhost:9091"
echo "  - Grafana:         http://localhost:3001 (admin/admin)"
echo "  - PostgreSQL:      localhost:5432"
echo "  - Redis:           localhost:6379"
echo ""
echo "🎉 Happy coding!"
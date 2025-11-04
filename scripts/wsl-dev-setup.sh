#!/bin/bash

# WSL Development Setup Script for Globepay
# This script helps set up and use the Globepay project in WSL

set -e

PROJECT_PATH="/mnt/c/Users/Leke/Globepay/globepay"

echo "🚀 Setting up Globepay development environment in WSL..."

# Check if we're in WSL
if ! grep -q microsoft /proc/version; then
    echo "❌ This script must be run in WSL"
    exit 1
fi

# Navigate to project directory
cd $PROJECT_PATH

echo "📂 Current directory: $(pwd)"

# Check prerequisites
echo "🔍 Checking prerequisites..."

if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed"
    echo "Please install Go in WSL:"
    echo "  sudo apt update"
    echo "  sudo apt install golang-go"
    exit 1
fi

if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed or not accessible"
    echo "Please ensure Docker Desktop is running and WSL integration is enabled"
    exit 1
fi

if ! command -v kubectl &> /dev/null; then
    echo "⚠️  kubectl is not installed"
    echo "To install kubectl:"
    echo "  curl -LO \"https://dl.k8s.io/release/v1.28.0/bin/linux/amd64/kubectl\""
    echo "  sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl"
fi

echo "✅ All prerequisites found"

# Setup backend
echo "🔨 Setting up backend..."
cd $PROJECT_PATH/backend
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init globepay
fi

echo "Downloading Go dependencies..."
go mod download

# Setup frontend
echo "⚛️  Setting up frontend..."
cd $PROJECT_PATH/frontend
if [ ! -f package.json ]; then
    echo "⚠️  package.json not found in frontend directory"
else
    # Check if npm is installed
    if ! command -v npm &> /dev/null; then
        echo "⚠️  npm is not installed"
        echo "To install Node.js and npm in WSL:"
        echo "  curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -"
        echo "  sudo apt-get install -y nodejs"
    else
        echo "Installing frontend dependencies..."
        npm install
    fi
fi

# Create .env files if they don't exist
echo "⚙️  Checking environment files..."
cd $PROJECT_PATH

if [ ! -f backend/.env ]; then
    echo "Creating backend .env file..."
    cp backend/.env.example backend/.env 2>/dev/null || echo "No .env.example found for backend"
fi

if [ ! -f frontend/.env ]; then
    echo "Creating frontend .env file..."
    cp frontend/.env.example frontend/.env 2>/dev/null || echo "No .env.example found for frontend"
fi

echo "✅ WSL development environment setup complete!"

echo ""
echo "📝 To start development:"
echo "  cd $PROJECT_PATH"
echo "  make dev-up     # Start development environment"
echo "  make dev-logs   # View logs"
echo "  make dev-down   # Stop development environment"
echo ""
echo "📝 To run tests:"
echo "  make test-backend   # Run backend tests"
echo "  make test-frontend  # Run frontend tests"
echo ""
echo "📝 To deploy:"
echo "  make deploy-dev     # Deploy to dev environment"
echo "  make deploy-staging # Deploy to staging environment"
echo "  make deploy-prod    # Deploy to production environment"
echo "  make deploy-all     # Deploy to all environments"
#!/bin/bash

set -e

echo "🔍 Verifying Globepay project setup..."

# Check if required tools are installed
echo "Checking required tools..."

# Check Git
if ! command -v git &> /dev/null
then
    echo "❌ Git is not installed"
    exit 1
else
    echo "✅ Git is installed"
fi

# Check Docker
if ! command -v docker &> /dev/null
then
    echo "❌ Docker is not installed"
    exit 1
else
    echo "✅ Docker is installed"
fi

# Check Docker Compose
if ! command -v docker-compose &> /dev/null
then
    echo "❌ Docker Compose is not installed"
    exit 1
else
    echo "✅ Docker Compose is installed"
fi

# Check Go
if ! command -v go &> /dev/null
then
    echo "❌ Go is not installed"
    exit 1
else
    echo "✅ Go is installed"
fi

# Check Node.js
if ! command -v node &> /dev/null
then
    echo "❌ Node.js is not installed"
    exit 1
else
    echo "✅ Node.js is installed"
fi

# Check if project structure exists
echo "Checking project structure..."

if [ ! -d "/mnt/c/Users/Leke/Globepay/globepay/backend" ]; then
    echo "❌ Backend directory not found"
    exit 1
else
    echo "✅ Backend directory exists"
fi

if [ ! -d "/mnt/c/Users/Leke/Globepay/globepay/frontend" ]; then
    echo "❌ Frontend directory not found"
    exit 1
else
    echo "✅ Frontend directory exists"
fi

if [ ! -d "/mnt/c/Users/Leke/Globepay/globepay/infrastructure" ]; then
    echo "❌ Infrastructure directory not found"
    exit 1
else
    echo "✅ Infrastructure directory exists"
fi

# Check backend setup
echo "Checking backend setup..."

if [ ! -f "/mnt/c/Users/Leke/Globepay/globepay/backend/go.mod" ]; then
    echo "❌ Go modules not initialized"
else
    echo "✅ Go modules initialized"
fi

# Check frontend setup
echo "Checking frontend setup..."

if [ ! -f "/mnt/c/Users/Leke/Globepay/globepay/frontend/package.json" ]; then
    echo "❌ Frontend dependencies not installed"
else
    echo "✅ Frontend dependencies installed"
fi

# Check documentation
echo "Checking documentation..."

if [ ! -f "/mnt/c/Users/Leke/Globepay/globepay/README.md" ]; then
    echo "❌ README.md not found"
else
    echo "✅ README.md exists"
fi

echo "🎉 Project setup verification completed successfully!"
echo ""
echo "Next steps:"
echo "1. Run './scripts/setup-dev-environment.sh' to set up your development environment"
echo "2. Run 'docker-compose up -d' to start all services"
echo "3. Access the application at http://localhost:3000"
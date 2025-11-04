# WSL Development Guide

This guide explains how to set up and use the Globepay project in Windows Subsystem for Linux (WSL).

## Prerequisites

1. Windows 10/11 with WSL2 installed
2. Ubuntu distribution installed in WSL
3. Docker Desktop for Windows with WSL2 integration enabled

## Setting Up WSL Development Environment

### 1. Install Required Tools in WSL

```bash
# Update package list
sudo apt update

# Install Go
sudo apt install golang-go

# Install other development tools
sudo apt install build-essential curl wget git
```

### 2. Run the Setup Script

```bash
# Navigate to the project directory
cd /mnt/c/Users/Leke/Globepay/globepay

# Make the setup script executable
chmod +x scripts/wsl-dev-setup.sh

# Run the setup script
./scripts/wsl-dev-setup.sh
```

## Using WSL for Development

### Starting the Development Environment

```bash
# Navigate to the project directory
cd /mnt/c/Users/Leke/Globepay/globepay

# Start all services
make dev-up

# View logs
make dev-logs

# Stop services
make dev-down
```

### Running Tests

```bash
# Run backend tests
make test-backend

# Run frontend tests
make test-frontend

# Run all tests
make test
```

### Building Docker Images

```bash
# Build backend Docker image
make build-backend

# Build frontend Docker image
make build-frontend

# Build all Docker images
make build
```

### Deploying to Environments

```bash
# Deploy to development
make deploy-dev

# Deploy to staging
make deploy-staging

# Deploy to production
make deploy-prod

# Deploy to all environments
make deploy-all
```

## Important Notes for WSL Development

### Path Translation

Windows paths are accessible in WSL through the `/mnt/` directory:
- Windows: `C:\Users\Leke\Globepay\globepay`
- WSL: `/mnt/c/Users/Leke/Globepay/globepay`

### Docker Integration

Ensure Docker Desktop is running and WSL2 integration is enabled:
1. Open Docker Desktop
2. Go to Settings > Resources > WSL Integration
3. Enable integration with your WSL distribution

### Command Chaining

In WSL, avoid using `&&` to chain commands. Instead, run commands sequentially:

```bash
# Instead of: go build && ./binary
# Use:
go build
./binary
```

### File Permissions

When creating scripts or executables in WSL, ensure they have proper permissions:

```bash
chmod +x script.sh
```

## Troubleshooting

### Docker Connection Issues

If you encounter Docker connection issues:
1. Ensure Docker Desktop is running
2. Restart Docker Desktop
3. Restart WSL: `wsl --shutdown` in PowerShell

### Go Module Issues

If you encounter Go module issues:
1. Ensure Git is installed: `sudo apt install git`
2. Check your GOPATH: `echo $GOPATH`
3. Verify Go version: `go version`

### Permission Denied Errors

If you encounter permission denied errors:
1. Check file permissions: `ls -la filename`
2. Fix permissions: `chmod +x filename`
3. Ensure you're running commands in the correct directory

## Useful WSL Commands

```bash
# List WSL distributions
wsl --list --verbose

# Shutdown WSL
wsl --shutdown

# Run a specific distribution
wsl -d Ubuntu

# Navigate to Windows drives
cd /mnt/c/Users/YourUsername/
```
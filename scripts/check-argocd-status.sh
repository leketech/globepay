#!/bin/bash

# Script to check the status of applications in ArgoCD
# Usage: ./scripts/check-argocd-status.sh

set -e

echo "Checking ArgoCD application status..."

# Check if argocd CLI is installed
if ! command -v argocd &> /dev/null; then
    echo "ArgoCD CLI is not installed. Installing..."
    # Install argocd CLI
    curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
    chmod +x /usr/local/bin/argocd
fi

# Check if we're logged in to ArgoCD
if ! argocd account list &> /dev/null; then
    echo "Not logged in to ArgoCD. Please log in first:"
    echo "argocd login <ARGOCD_SERVER> --username <USERNAME> --password <PASSWORD>"
    exit 1
fi

# List all applications
echo "=== ArgoCD Applications ==="
argocd app list

# Check specific applications
echo -e "\n=== Backend Application Status ==="
argocd app get globepay-backend-prod-helm

echo -e "\n=== Frontend Application Status ==="
argocd app get globepay-frontend-prod-helm

# Check for any out-of-sync applications
echo -e "\n=== Out-of-Sync Applications ==="
argocd app list | grep -v "Synced"

echo "ArgoCD status check completed!"
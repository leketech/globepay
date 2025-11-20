#!/bin/bash

# Script to sync ArgoCD applications
# Usage: ./scripts/sync-argocd-apps.sh

set -e

echo "Syncing ArgoCD applications..."

# Check if argocd CLI is installed
if ! command -v argocd &> /dev/null; then
    echo "ArgoCD CLI is not installed. Please install it first."
    exit 1
fi

# Check if we're logged in to ArgoCD
if ! argocd account list &> /dev/null; then
    echo "Not logged in to ArgoCD. Please log in first:"
    echo "argocd login <ARGOCD_SERVER> --username <USERNAME> --password <PASSWORD>"
    exit 1
fi

# Sync backend application
echo "Syncing backend application..."
argocd app sync globepay-backend-prod-helm --prune

# Sync frontend application
echo "Syncing frontend application..."
argocd app sync globepay-frontend-prod-helm --prune

# Wait for sync to complete
echo "Waiting for sync to complete..."
argocd app wait globepay-backend-prod-helm --health
argocd app wait globepay-frontend-prod-helm --health

echo "ArgoCD applications synced successfully!"
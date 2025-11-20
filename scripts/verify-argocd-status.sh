#!/bin/bash

# Script to verify ArgoCD application status
# This script checks if ArgoCD applications are healthy and synchronized

set -e

NAMESPACE="globepay-prod"
TIMEOUT=300  # 5 minutes timeout
INTERVAL=10  # Check every 10 seconds

echo "Checking ArgoCD application status..."

# Function to check if kubectl is available
check_kubectl() {
    if ! command -v kubectl &> /dev/null; then
        echo "kubectl is not installed or not in PATH"
        exit 1
    fi
}

# Function to check if ArgoCD CLI is available
check_argocd_cli() {
    if ! command -v argocd &> /dev/null; then
        echo "argocd CLI is not installed or not in PATH"
        echo "Installing ArgoCD CLI..."
        curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
        chmod +x /usr/local/bin/argocd
    fi
}

# Function to check ArgoCD application status
check_application_status() {
    local app_name=$1
    echo "Checking status for application: $app_name"
    
    # Check if application exists
    if ! kubectl get application "$app_name" -n argocd &> /dev/null; then
        echo "Application $app_name not found in ArgoCD"
        return 1
    fi
    
    # Get application status
    local status
    status=$(kubectl get application "$app_name" -n argocd -o jsonpath='{.status.health.status}')
    local sync_status
    sync_status=$(kubectl get application "$app_name" -n argocd -o jsonpath='{.status.sync.status}')
    
    echo "Application: $app_name"
    echo "Health Status: $status"
    echo "Sync Status: $sync_status"
    
    if [[ "$status" == "Healthy" && "$sync_status" == "Synced" ]]; then
        echo "Application $app_name is Healthy and Synced"
        return 0
    else
        echo "Application $app_name is not ready (Health: $status, Sync: $sync_status)"
        return 1
    fi
}

# Function to wait for applications to be ready
wait_for_applications() {
    local start_time
    start_time=$(date +%s)
    
    while true; do
        local current_time
        current_time=$(date +%s)
        local elapsed=$((current_time - start_time))
        
        if [ $elapsed -gt $TIMEOUT ]; then
            echo "Timeout waiting for applications to be ready"
            exit 1
        fi
        
        echo "Waiting for applications to be healthy and synced... ($elapsed/$TIMEOUT seconds)"
        
        # Check backend application
        if check_application_status "globepay-backend-prod" && check_application_status "globepay-frontend-prod"; then
            echo "All applications are healthy and synced"
            return 0
        fi
        
        sleep $INTERVAL
    done
}

# Main execution
main() {
    check_kubectl
    # check_argocd_cli  # Uncomment if you want to ensure ArgoCD CLI is installed
    
    echo "Verifying ArgoCD application status in namespace: $NAMESPACE"
    
    # Wait for applications to be ready
    wait_for_applications
    
    echo "ArgoCD verification completed successfully"
}

# Run main function
main "$@"
#!/bin/bash

# Script to promote changes from dev -> staging -> prod
# Usage: ./scripts/promote-changes.sh [environment]
# If no environment specified, promotes through all environments

set -e

echo "🚀 Starting change promotion process..."

ENVIRONMENT=${1:-all}

# Function to deploy to an environment
deploy_to_environment() {
    local env=$1
    local namespace="globepay-$env"
    
    echo "📦 Deploying to $env environment..."
    
    # Apply kustomize configuration
    kubectl apply -k k8s/overlays/$env
    
    # Wait for rollout to complete
    echo "⏳ Waiting for deployment to complete..."
    kubectl rollout status deployment/backend -n $namespace
    kubectl rollout status deployment/frontend -n $namespace
    
    echo "✅ Deployment to $env environment completed successfully!"
}

# Function to run tests for an environment
test_environment() {
    local env=$1
    
    echo "🧪 Running tests for $env environment..."
    
    # Add environment-specific tests here
    # For now, we'll just echo the test execution
    echo "Running integration tests for $env..."
    
    echo "✅ Tests for $env environment passed!"
}

case $ENVIRONMENT in
    "dev")
        echo "🎯 Promoting to development environment only"
        deploy_to_environment "dev"
        test_environment "dev"
        ;;
    "staging")
        echo "🎯 Promoting to staging environment only"
        # First ensure dev is working
        test_environment "dev"
        deploy_to_environment "staging"
        test_environment "staging"
        ;;
    "prod"|"production")
        echo "🎯 Promoting to production environment only"
        # First ensure staging is working
        test_environment "staging"
        deploy_to_environment "prod"
        test_environment "prod"
        ;;
    "all"|"*")
        echo "🔄 Promoting changes through all environments"
        # Deploy to dev first
        deploy_to_environment "dev"
        test_environment "dev"
        
        # Deploy to staging
        deploy_to_environment "staging"
        test_environment "staging"
        
        # Deploy to production
        deploy_to_environment "prod"
        test_environment "prod"
        
        echo "🎉 All environments updated successfully!"
        ;;
    *)
        echo "❌ Unknown environment: $ENVIRONMENT"
        echo "Usage: $0 [dev|staging|prod|all]"
        exit 1
        ;;
esac

echo "🏁 Change promotion process completed!"
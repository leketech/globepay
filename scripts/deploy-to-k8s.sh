#!/bin/bash

# Script to build, push, and deploy the application to Kubernetes
# Usage: ./scripts/deploy-to-k8s.sh <environment> <tag>

set -e

if [ $# -lt 2 ]; then
    echo "Usage: $0 <environment> <tag>"
    echo "Example: $0 prod v1.0.0"
    exit 1
fi

ENVIRONMENT=$1
TAG=$2

echo "Deploying to environment: $ENVIRONMENT with tag: $TAG"

# Define variables
ECR_REGISTRY="907849381252.dkr.ecr.us-east-1.amazonaws.com"
BACKEND_IMAGE="$ECR_REGISTRY/globepay-backend:$TAG"
FRONTEND_IMAGE="$ECR_REGISTRY/globepay-frontend:$TAG"

# Build backend Docker image
echo "Building backend Docker image..."
docker build -t $BACKEND_IMAGE ./backend

# Build frontend Docker image
echo "Building frontend Docker image..."
docker build -t $FRONTEND_IMAGE ./frontend

# Push images to ECR
echo "Pushing images to ECR..."
docker push $BACKEND_IMAGE
docker push $FRONTEND_IMAGE

# Update Kubernetes manifests with new image tags
echo "Updating Kubernetes manifests..."
./scripts/update-image-tags.sh $ENVIRONMENT $TAG $TAG

# Update secret patch with actual values
echo "Updating secret patch..."
./scripts/update-secret-patch.sh $ENVIRONMENT

# Deploy to Kubernetes
echo "Deploying to Kubernetes..."
case $ENVIRONMENT in
    "prod")
        kubectl apply -k k8s/overlays/prod
        kubectl apply -k k8s/overlays/frontend-prod
        ;;
    "staging")
        kubectl apply -k k8s/overlays/staging
        ;;
    "dev")
        kubectl apply -k k8s/overlays/dev
        ;;
    *)
        echo "Unknown environment: $ENVIRONMENT"
        exit 1
        ;;
esac

echo "Deployment completed successfully!"
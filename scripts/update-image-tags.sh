#!/bin/bash

# Script to update image tags in Kubernetes deployment patches
# Usage: ./scripts/update-image-tags.sh <environment> <backend-tag> <frontend-tag>

set -e

if [ $# -lt 3 ]; then
    echo "Usage: $0 <environment> <backend-tag> <frontend-tag>"
    echo "Example: $0 prod v1.0.0 v1.0.0"
    exit 1
fi

ENVIRONMENT=$1
BACKEND_TAG=$2
FRONTEND_TAG=$3

echo "Updating image tags for environment: $ENVIRONMENT"
echo "Backend tag: $BACKEND_TAG"
echo "Frontend tag: $FRONTEND_TAG"

# Define ECR repository
ECR_REGISTRY="907849381252.dkr.ecr.us-east-1.amazonaws.com"
BACKEND_IMAGE="$ECR_REGISTRY/globepay-backend:$BACKEND_TAG"
FRONTEND_IMAGE="$ECR_REGISTRY/globepay-frontend:$FRONTEND_TAG"

# Update backend deployment patch based on environment
case $ENVIRONMENT in
    "prod")
        BACKEND_PATCH_FILE="k8s/overlays/prod/patches/backend-deployment.yaml"
        FRONTEND_PATCH_FILE="k8s/overlays/frontend-prod/patches/frontend-deployment.yaml"
        ;;
    "staging")
        BACKEND_PATCH_FILE="k8s/overlays/staging/patches/backend-deployment.yaml"
        FRONTEND_PATCH_FILE="k8s/overlays/staging/patches/frontend-deployment.yaml"
        ;;
    "dev")
        BACKEND_PATCH_FILE="k8s/overlays/dev/patches/backend-deployment.yaml"
        FRONTEND_PATCH_FILE="k8s/overlays/dev/patches/frontend-deployment.yaml"
        ;;
    *)
        echo "Unknown environment: $ENVIRONMENT"
        exit 1
        ;;
esac

# Update backend image tag
if [ -f "$BACKEND_PATCH_FILE" ]; then
    echo "Updating backend image in $BACKEND_PATCH_FILE"
    sed -i "s|image: .*globepay-backend:.*|image: $BACKEND_IMAGE|" "$BACKEND_PATCH_FILE"
else
    echo "Warning: $BACKEND_PATCH_FILE not found"
fi

# Update frontend image tag
if [ -f "$FRONTEND_PATCH_FILE" ]; then
    echo "Updating frontend image in $FRONTEND_PATCH_FILE"
    sed -i "s|image: .*globepay-frontend:.*|image: $FRONTEND_IMAGE|" "$FRONTEND_PATCH_FILE"
else
    echo "Warning: $FRONTEND_PATCH_FILE not found"
fi

echo "Image tags updated successfully!"
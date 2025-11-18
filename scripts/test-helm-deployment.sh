#!/bin/bash

# Script to test Helm deployment locally
# Usage: ./scripts/test-helm-deployment.sh

set -e

echo "Testing Helm deployment..."

# Check if Helm is installed
if ! command -v helm &> /dev/null; then
    echo "Helm is not installed. Please install Helm first."
    exit 1
fi

# Validate the Helm chart
echo "Validating Helm chart..."
helm lint ./helm/globepay

# Template the Helm chart to see what would be deployed
echo "Templating Helm chart..."
helm template globepay-test ./helm/globepay \
  -f ./helm/globepay/values-prod.yaml \
  --set secrets.dbPassword="test-password" \
  --set secrets.jwtSecret="test-secret" \
  --namespace globepay-test

echo "Helm deployment test completed successfully!"
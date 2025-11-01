#!/bin/bash

# Destroy Terraform environment

set -e

ENVIRONMENT=$1

if [ -z "$ENVIRONMENT" ]; then
  echo "Usage: $0 <environment>"
  echo "Example: $0 dev"
  exit 1
fi

echo "Destroying Terraform environment: $ENVIRONMENT"

# Validate environment
if [ "$ENVIRONMENT" != "dev" ] && [ "$ENVIRONMENT" != "staging" ] && [ "$ENVIRONMENT" != "prod" ]; then
  echo "Invalid environment. Must be one of: dev, staging, prod"
  exit 1
fi

# Confirm destruction
echo "WARNING: This will destroy all resources in the $ENVIRONMENT environment!"
read -p "Are you sure you want to continue? (yes/no): " CONFIRM

if [ "$CONFIRM" != "yes" ]; then
  echo "Destruction cancelled."
  exit 0
fi

# Destroy environment
cd ../environments/$ENVIRONMENT
terraform destroy -auto-approve

echo "Environment $ENVIRONMENT destroyed successfully!"
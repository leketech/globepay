#!/bin/bash

# Initialize Terraform for all environments

set -e

echo "Initializing Terraform for all environments..."

# Initialize development environment
echo "Initializing development environment..."
cd ../environments/dev
terraform init

cd ..

# Initialize staging environment
echo "Initializing staging environment..."
cd ../environments/staging
terraform init

cd ..

# Initialize production environment
echo "Initializing production environment..."
cd ../environments/prod
terraform init

echo "Terraform initialization complete for all environments!"
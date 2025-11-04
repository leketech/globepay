#!/bin/bash

# Staging Deployment Script

echo "Deploying to Staging Environment..."

# Load environment variables
if [ -f .env.staging ]; then
    export $(cat .env.staging | xargs)
fi

# Build and start services
docker-compose -f docker-compose.staging.yml up -d

echo "Staging deployment completed!"
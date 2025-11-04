#!/bin/bash

# Production Deployment Script

echo "Deploying to Production Environment..."

# Load environment variables
if [ -f .env.prod ]; then
    export $(cat .env.prod | xargs)
fi

# Build and start services
docker-compose -f docker-compose.prod.yml up -d

echo "Production deployment completed!"
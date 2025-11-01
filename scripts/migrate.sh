#!/bin/bash

set -e

echo "Running database migrations..."

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null
then
    echo "docker-compose could not be found, trying docker compose"
    if ! command -v docker &> /dev/null
    then
        echo "docker could not be found"
        exit 1
    fi
    DOCKER_COMPOSE="docker compose"
else
    DOCKER_COMPOSE="docker-compose"
fi

# Run migrations
$DOCKER_COMPOSE -f docker-compose.migration.yml up --abort-on-container-exit

echo "Database migrations completed!"
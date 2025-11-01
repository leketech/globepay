#!/bin/bash

set -e

echo "Creating database backup..."

# Create backup directory
BACKUP_DIR="./backups"
mkdir -p $BACKUP_DIR

# Get current timestamp
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

# Create backup
docker-compose exec postgres pg_dump -U postgres globepay > $BACKUP_DIR/globepay_backup_$TIMESTAMP.sql

echo "Backup created: $BACKUP_DIR/globepay_backup_$TIMESTAMP.sql"
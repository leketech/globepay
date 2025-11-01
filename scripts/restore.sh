#!/bin/bash

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <backup-file>"
    exit 1
fi

BACKUP_FILE=$1

if [ ! -f "$BACKUP_FILE" ]; then
    echo "Backup file not found: $BACKUP_FILE"
    exit 1
fi

echo "Restoring database from $BACKUP_FILE..."

# Stop backend service to prevent conflicts
docker-compose stop backend

# Restore database
docker-compose exec -T postgres psql -U postgres globepay < $BACKUP_FILE

# Start backend service
docker-compose start backend

echo "Database restored successfully!"
#!/bin/bash

# Script to update configmap patch with actual secret values
# Usage: ./scripts/update-configmap-patch.sh <environment>

set -e

if [ $# -lt 1 ]; then
    echo "Usage: $0 <environment>"
    echo "Example: $0 prod"
    exit 1
fi

ENVIRONMENT=$1
ENV_FILE="k8s/overlays/$ENVIRONMENT/.env.secret"
CONFIGMAP_PATCH_FILE="k8s/overlays/$ENVIRONMENT/patches/configmap-patch.yaml"

echo "Updating configmap patch for environment: $ENVIRONMENT"

# Check if env file exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: Environment file $ENV_FILE not found"
    exit 1
fi

# Check if configmap patch file exists
if [ ! -f "$CONFIGMAP_PATCH_FILE" ]; then
    echo "Error: ConfigMap patch file $CONFIGMAP_PATCH_FILE not found"
    exit 1
fi

# Create a temporary file
TEMP_FILE=$(mktemp)

# Copy the original file to temporary file
cp "$CONFIGMAP_PATCH_FILE" "$TEMP_FILE"

# Read each line from the env file and substitute variables
while IFS= read -r line; do
    if [[ $line == *"="* ]]; then
        key="${line%%=*}"
        value="${line#*=}"
        echo "Substituting $key with actual value"
        sed -i "s|\${$key}|$value|g" "$TEMP_FILE"
    fi
done < "$ENV_FILE"

# Move the updated file back
mv "$TEMP_FILE" "$CONFIGMAP_PATCH_FILE"

echo "ConfigMap patch updated successfully!"
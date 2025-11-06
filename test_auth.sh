#!/bin/bash

# Test the auth endpoints from within the cluster
echo "Testing registration..."
/usr/local/bin/kubectl exec -n globepay-prod frontend-7b69d6d96-l2mkf -- curl -v -X POST \
  -H "Content-Type: application/json" \
  -d '{"email": "testuser@example.com", "password": "securepassword123", "firstName": "Test", "lastName": "User"}' \
  http://backend/api/v1/auth/register

echo -e "\nTesting login..."
/usr/local/bin/kubectl exec -n globepay-prod frontend-7b69d6d96-l2mkf -- curl -v -X POST \
  -H "Content-Type: application/json" \
  -d '{"email": "testuser@example.com", "password": "securepassword123"}' \
  http://backend/api/v1/auth/login
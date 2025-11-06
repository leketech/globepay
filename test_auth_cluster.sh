#!/bin/bash

# Test the auth endpoints from within the cluster
echo "Testing registration..."
/usr/local/bin/kubectl exec -n globepay-prod frontend-7b69d6d96-2qlmk -- curl -v -X POST \
  -H "Content-Type: application/json" \
  -d '{"email": "testuser3@example.com", "password": "securepassword123", "firstName": "Test", "lastName": "User"}' \
  http://backend/api/v1/auth/register

echo -e "\nTesting login..."
/usr/local/bin/kubectl exec -n globepay-prod frontend-7b69d6d96-2qlmk -- curl -v -X POST \
  -H "Content-Type: application/json" \
  -d '{"email": "testuser3@example.com", "password": "securepassword123"}' \
  http://backend/api/v1/auth/login
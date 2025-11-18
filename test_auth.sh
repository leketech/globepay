#!/bin/bash

# Test script for login and signup functionality
API_URL="https://api.globepay.space"

echo "Testing signup endpoint..."
curl -k -X POST ${API_URL}/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "firstName": "Test",
    "lastName": "User"
  }'

echo -e "\n\nTesting login endpoint..."
curl -k -X POST ${API_URL}/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'

echo -e "\n"
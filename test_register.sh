#!/bin/bash

# Test registration
echo "Testing registration..."
kubectl -n globepay-prod run test-register --rm -it --image=curlimages/curl -- \
  curl -s -X POST -H "Content-Type: application/json" \
  -d '{"email":"testuser@globepay.space","password":"TestPassword123!","firstName":"Test","lastName":"User"}' \
  http://10.0.10.200:8080/api/v1/auth/register
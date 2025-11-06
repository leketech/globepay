#!/bin/sh
curl -X POST http://backend:80/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"newuser@example.com","password":"password123","firstName":"New","lastName":"User"}'
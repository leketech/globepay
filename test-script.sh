#!/bin/sh
curl -X POST http://backend:80/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
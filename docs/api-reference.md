# API Reference Documentation

This document provides detailed reference information for all Globepay API endpoints.

## Base URL

```
https://api.globepay.com/api/v1
```

For local development:
```
http://localhost:8080/api/v1
```

## Authentication

All API endpoints (except authentication endpoints) require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

### Token Expiration

- **Access Token**: 15 minutes
- **Refresh Token**: 7 days

## Rate Limiting

API endpoints are rate-limited to prevent abuse:
- **100 requests per minute** per IP address
- **1000 requests per hour** per authenticated user

Exceeding these limits will result in a `429 Too Many Requests` response.

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

### HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 422 | Unprocessable Entity |
| 429 | Too Many Requests |
| 500 | Internal Server Error |
| 503 | Service Unavailable |

## Authentication Endpoints

### POST /auth/login

Authenticate a user with email and password.

**Request:**
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refreshToken": "refresh-token-here",
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "kycStatus": "verified",
    "accountStatus": "active"
  }
}
```

**Response Codes:**
- `200`: Successful login
- `400`: Invalid request body
- `401`: Invalid credentials
- `429`: Rate limit exceeded

### POST /auth/register

Register a new user.

**Request:**
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890",
  "dateOfBirth": "1990-01-01",
  "country": "US"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refreshToken": "refresh-token-here",
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "kycStatus": "pending",
    "accountStatus": "active"
  }
}
```

**Response Codes:**
- `201`: User created successfully
- `400`: Invalid request body
- `409`: Email already exists
- `429`: Rate limit exceeded

### POST /auth/refresh

Refresh JWT token using refresh token.

**Request:**
```http
POST /api/v1/auth/refresh
Content-Type: application/json

{
  "refreshToken": "refresh-token-here"
}
```

**Response:**
```json
{
  "token": "new-jwt-token-here",
  "refreshToken": "new-refresh-token-here"
}
```

**Response Codes:**
- `200`: Token refreshed successfully
- `400`: Invalid refresh token
- `401`: Expired refresh token

### POST /auth/forgot-password

Initiate password reset process.

**Request:**
```http
POST /api/v1/auth/forgot-password
Content-Type: application/json

{
  "email": "user@example.com"
}
```

**Response:**
```json
{
  "message": "Password reset email sent"
}
```

**Response Codes:**
- `200`: Reset email sent
- `400`: Invalid email
- `404`: User not found
- `429`: Rate limit exceeded

### POST /auth/reset-password

Reset user password using token.

**Request:**
``http
POST /api/v1/auth/reset-password
Content-Type: application/json

{
  "token": "reset-token-here",
  "password": "newpassword123"
}
```

**Response:**
``json
{
  "message": "Password reset successfully"
}
```

**Response Codes:**
- `200`: Password reset successfully
- `400`: Invalid token or password
- `404`: Token not found

## User Endpoints

### GET /user/profile

Get authenticated user's profile information.

**Request:**
```http
GET /api/v1/user/profile
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890",
  "dateOfBirth": "1990-01-01",
  "country": "US",
  "kycStatus": "verified",
  "accountStatus": "active",
  "emailVerified": true,
  "phoneVerified": true,
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Profile retrieved successfully
- `401`: Unauthorized

### PUT /user/profile

Update user profile information.

**Request:**
```http
PUT /api/v1/user/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890"
}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890",
  "dateOfBirth": "1990-01-01",
  "country": "US",
  "kycStatus": "verified",
  "accountStatus": "active",
  "emailVerified": true,
  "phoneVerified": true,
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Profile updated successfully
- `400`: Invalid request body
- `401`: Unauthorized

### GET /user/accounts

Get user's financial accounts.

**Request:**
```http
GET /api/v1/user/accounts
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "userId": "123e4567-e89b-12d3-a456-426614174000",
    "currency": "USD",
    "balance": 1000.50,
    "accountNumber": "ACC123456789",
    "accountType": "checking",
    "status": "active",
    "createdAt": "2023-01-01T00:00:00Z",
    "updatedAt": "2023-01-01T00:00:00Z"
  }
]
```

**Response Codes:**
- `200`: Accounts retrieved successfully
- `401`: Unauthorized

### POST /user/accounts

Create a new financial account.

**Request:**
```http
POST /api/v1/user/accounts
Authorization: Bearer <token>
Content-Type: application/json

{
  "currency": "EUR"
}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "currency": "EUR",
  "balance": 0.00,
  "accountNumber": "ACC987654321",
  "accountType": "checking",
  "status": "active",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `201`: Account created successfully
- `400`: Invalid request body
- `401`: Unauthorized
- `409`: Account already exists for currency

## Transfer Endpoints

### GET /transfers

Get user's transfers with pagination.

**Request:**
```http
GET /api/v1/transfers?page=1&limit=10
Authorization: Bearer <token>
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)

**Response:**
``json
{
  "transfers": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "userId": "123e4567-e89b-12d3-a456-426614174000",
      "recipientName": "Jane Smith",
      "recipientCountry": "GB",
      "sourceAmount": 100.00,
      "destAmount": 85.50,
      "sourceCurrency": "USD",
      "destCurrency": "GBP",
      "exchangeRate": 0.855,
      "fee": 2.50,
      "purpose": "family_support",
      "status": "completed",
      "referenceNumber": "REF123456",
      "estimatedArrival": "2023-01-02T00:00:00Z",
      "processedAt": "2023-01-01T00:00:00Z",
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

**Response Codes:**
- `200`: Transfers retrieved successfully
- `400`: Invalid query parameters
- `401`: Unauthorized

### GET /transfers/{id}

Get a specific transfer by ID.

**Request:**
```http
GET /api/v1/transfers/123e4567-e89b-12d3-a456-426614174000
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "recipientName": "Jane Smith",
  "recipientEmail": "jane@example.com",
  "recipientCountry": "GB",
  "recipientBankName": "Bank of England",
  "recipientAccountNumber": "12345678",
  "recipientSwiftCode": "BOEGB22",
  "sourceAmount": 100.00,
  "destAmount": 85.50,
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "exchangeRate": 0.855,
  "fee": 2.50,
  "purpose": "family_support",
  "status": "completed",
  "referenceNumber": "REF123456",
  "estimatedArrival": "2023-01-02T00:00:00Z",
  "processedAt": "2023-01-01T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Transfer retrieved successfully
- `401`: Unauthorized
- `403`: Forbidden (not owner of transfer)
- `404`: Transfer not found

### POST /transfers

Create a new money transfer.

**Request:**
```http
POST /api/v1/transfers
Authorization: Bearer <token>
Content-Type: application/json

{
  "recipientName": "Jane Smith",
  "recipientEmail": "jane@example.com",
  "recipientCountry": "GB",
  "recipientBankName": "Bank of England",
  "recipientAccountNumber": "12345678",
  "recipientSwiftCode": "BOEGB22",
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "sourceAmount": 100.00,
  "purpose": "family_support"
}
```

**Response:**
``json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "recipientName": "Jane Smith",
  "recipientEmail": "jane@example.com",
  "recipientCountry": "GB",
  "recipientBankName": "Bank of England",
  "recipientAccountNumber": "12345678",
  "recipientSwiftCode": "BOEGB22",
  "sourceAmount": 100.00,
  "destAmount": 85.50,
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "exchangeRate": 0.855,
  "fee": 2.50,
  "purpose": "family_support",
  "status": "pending",
  "referenceNumber": "REF123456",
  "estimatedArrival": "2023-01-02T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `201`: Transfer created successfully
- `400`: Invalid request body
- `401`: Unauthorized
- `402`: Insufficient funds
- `422`: Validation errors

### POST /transfers/{id}/cancel

Cancel a pending transfer.

**Request:**
```http
POST /api/v1/transfers/123e4567-e89b-12d3-a456-426614174000/cancel
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Transfer cancelled successfully"
}
```

**Response Codes:**
- `200`: Transfer cancelled successfully
- `401`: Unauthorized
- `403`: Forbidden (not owner of transfer)
- `404`: Transfer not found
- `409`: Transfer cannot be cancelled (wrong status)

### GET /transfers/rates

Get exchange rates for currency conversion.

**Request:**
```http
GET /api/v1/transfers/rates?from=USD&to=GBP&amount=100
Authorization: Bearer <token>
```

**Query Parameters:**
- `from` (required): Source currency code
- `to` (required): Destination currency code
- `amount` (required): Amount to convert

**Response:**
```json
{
  "fromCurrency": "USD",
  "toCurrency": "GBP",
  "rate": 0.855,
  "fee": 2.50,
  "amount": 100.00,
  "convertedAmount": 85.50,
  "totalFee": 2.50,
  "timestamp": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Exchange rate retrieved successfully
- `400`: Invalid query parameters
- `401`: Unauthorized

## Transaction Endpoints

### GET /transactions

Get user's transactions with pagination.

**Request:**
```http
GET /api/v1/transactions?page=1&limit=10
Authorization: Bearer <token>
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)

**Response:**
```json
{
  "transactions": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "userId": "123e4567-e89b-12d3-a456-426614174000",
      "accountId": "123e4567-e89b-12d3-a456-426614174000",
      "transferId": "123e4567-e89b-12d3-a456-426614174000",
      "type": "TRANSFER",
      "status": "completed",
      "amount": 100.00,
      "currency": "USD",
      "fee": 2.50,
      "description": "Transfer to Jane Smith",
      "referenceNumber": "REF123456",
      "processedAt": "2023-01-01T00:00:00Z",
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

**Response Codes:**
- `200`: Transactions retrieved successfully
- `400`: Invalid query parameters
- `401`: Unauthorized

### GET /transactions/{id}

Get a specific transaction by ID.

**Request:**
```http
GET /api/v1/transactions/123e4567-e89b-12d3-a456-426614174000
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "accountId": "123e4567-e89b-12d3-a456-426614174000",
  "transferId": "123e4567-e89b-12d3-a456-426614174000",
  "type": "TRANSFER",
  "status": "completed",
  "amount": 100.00,
  "currency": "USD",
  "fee": 2.50,
  "description": "Transfer to Jane Smith",
  "referenceNumber": "REF123456",
  "processedAt": "2023-01-01T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Transaction retrieved successfully
- `401`: Unauthorized
- `403`: Forbidden (not owner of transaction)
- `404`: Transaction not found

## Beneficiary Endpoints

### GET /beneficiaries

Get user's saved beneficiaries.

**Request:**
```http
GET /api/v1/beneficiaries
Authorization: Bearer <token>
```

**Response:**
``json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "userId": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Jane Smith",
    "country": "GB",
    "bankName": "Bank of England",
    "accountNumber": "12345678",
    "swiftCode": "BOEGB22",
    "createdAt": "2023-01-01T00:00:00Z",
    "updatedAt": "2023-01-01T00:00:00Z"
  }
]
```

**Response Codes:**
- `200`: Beneficiaries retrieved successfully
- `401`: Unauthorized

### POST /beneficiaries

Create a new beneficiary.

**Request:**
```http
POST /api/v1/beneficiaries
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Jane Smith",
  "country": "GB",
  "bankName": "Bank of England",
  "accountNumber": "12345678",
  "swiftCode": "BOEGB22"
}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Jane Smith",
  "country": "GB",
  "bankName": "Bank of England",
  "accountNumber": "12345678",
  "swiftCode": "BOEGB22",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `201`: Beneficiary created successfully
- `400`: Invalid request body
- `401`: Unauthorized
- `409`: Beneficiary already exists

### PUT /beneficiaries/{id}

Update a beneficiary.

**Request:**
```http
PUT /api/v1/beneficiaries/123e4567-e89b-12d3-a456-426614174000
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Jane Smith",
  "country": "GB",
  "bankName": "Bank of England",
  "accountNumber": "12345678",
  "swiftCode": "BOEGB22"
}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Jane Smith",
  "country": "GB",
  "bankName": "Bank of England",
  "accountNumber": "12345678",
  "swiftCode": "BOEGB22",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Beneficiary updated successfully
- `400`: Invalid request body
- `401`: Unauthorized
- `403`: Forbidden (not owner of beneficiary)
- `404`: Beneficiary not found

### DELETE /beneficiaries/{id}

Delete a beneficiary.

**Request:**
```http
DELETE /api/v1/beneficiaries/123e4567-e89b-12d3-a456-426614174000
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Beneficiary deleted successfully"
}
```

**Response Codes:**
- `200`: Beneficiary deleted successfully
- `401`: Unauthorized
- `403`: Forbidden (not owner of beneficiary)
- `404`: Beneficiary not found

## Health Check Endpoints

### GET /health

Basic health check endpoint.

**Request:**
```http
GET /api/v1/health
```

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Service is healthy

### GET /ready

Readiness check endpoint.

**Request:**
```http
GET /api/v1/ready
```

**Response:**
``json
{
  "status": "ready",
  "database": "connected",
  "cache": "connected",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

**Response Codes:**
- `200`: Service is ready
- `503`: Service is not ready

## Webhooks

### POST /webhooks/transfer-status

Webhook endpoint for transfer status updates.

**Request:**
```http
POST /api/v1/webhooks/transfer-status
Content-Type: application/json
X-Signature: sha256-hex-signature

{
  "eventType": "transfer.status.updated",
  "transferId": "123e4567-e89b-12d3-a456-426614174000",
  "status": "completed",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

**Response:**
```json
{
  "message": "Webhook received successfully"
}
```

**Response Codes:**
- `200`: Webhook processed successfully
- `400`: Invalid webhook payload
- `401`: Invalid signature

## API Versioning

The API uses URL versioning with `v1` as the current version. Future versions will be released as `v2`, `v3`, etc.

## SDKs

Official SDKs are available for:
- JavaScript/Node.js
- Python
- Go
- Java

Check the [SDK repository](https://github.com/your-org/globepay-sdks) for more information.

This API reference documentation provides detailed information about all available endpoints, request/response formats, and error handling for the Globepay API.
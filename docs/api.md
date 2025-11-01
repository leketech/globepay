# API Documentation

This document provides detailed information about the Globepay API endpoints.

## Authentication

All API endpoints (except authentication endpoints) require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

## Base URL

```
http://localhost:8080/api/v1
```

In production, this will be:
```
https://api.globepay.com/api/v1
```

## Authentication Endpoints

### POST /auth/login

Login with email and password.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "123",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe"
  }
}
```

### POST /auth/register

Register a new user.

**Request:**
```json
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
  "user": {
    "id": "123",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe"
  }
}
```

### POST /auth/refresh

Refresh JWT token.

**Request:**
```json
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

## User Endpoints

### GET /user/profile

Get user profile information.

**Response:**
```json
{
  "id": "123",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890",
  "dateOfBirth": "1990-01-01",
  "country": "US",
  "kycStatus": "verified",
  "accountStatus": "active",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

### PUT /user/profile

Update user profile information.

**Request:**
```json
{
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890"
}
```

**Response:**
```json
{
  "id": "123",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "phoneNumber": "+1234567890",
  "dateOfBirth": "1990-01-01",
  "country": "US",
  "kycStatus": "verified",
  "accountStatus": "active",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

### GET /user/accounts

Get user accounts.

**Response:**
```json
[
  {
    "id": "456",
    "userId": "123",
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

### POST /user/accounts

Create a new account.

**Request:**
```json
{
  "currency": "EUR"
}
```

**Response:**
```json
{
  "id": "789",
  "userId": "123",
  "currency": "EUR",
  "balance": 0.00,
  "accountNumber": "ACC987654321",
  "accountType": "checking",
  "status": "active",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

## Transfer Endpoints

### GET /transfers

Get user transfers with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "transfers": [
    {
      "id": "123",
      "userId": "456",
      "recipientName": "Jane Smith",
      "recipientCountry": "GB",
      "sourceAmount": 100.00,
      "destAmount": 85.50,
      "sourceCurrency": "USD",
      "destCurrency": "GBP",
      "exchangeRate": 0.855,
      "fee": 2.50,
      "status": "completed",
      "estimatedArrival": "2023-01-02T00:00:00Z",
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
    }
  ],
  "total": 1
}
```

### GET /transfers/{id}

Get a specific transfer by ID.

**Response:**
```json
{
  "id": "123",
  "userId": "456",
  "recipientName": "Jane Smith",
  "recipientCountry": "GB",
  "sourceAmount": 100.00,
  "destAmount": 85.50,
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "exchangeRate": 0.855,
  "fee": 2.50,
  "status": "completed",
  "estimatedArrival": "2023-01-02T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

### POST /transfers

Create a new transfer.

**Request:**
```json
{
  "recipientName": "Jane Smith",
  "recipientEmail": "jane@example.com",
  "recipientCountry": "GB",
  "recipientBankName": "Bank of England",
  "recipientAccountNo": "12345678",
  "recipientSwiftCode": "BOEGB22",
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "sourceAmount": 100.00,
  "purpose": "family_support"
}
```

**Response:**
```json
{
  "id": "123",
  "userId": "456",
  "recipientName": "Jane Smith",
  "recipientCountry": "GB",
  "sourceAmount": 100.00,
  "destAmount": 85.50,
  "sourceCurrency": "USD",
  "destCurrency": "GBP",
  "exchangeRate": 0.855,
  "fee": 2.50,
  "status": "pending",
  "estimatedArrival": "2023-01-02T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

### GET /transfers/rates

Get exchange rates.

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
  "timestamp": "2023-01-01T00:00:00Z"
}
```

## Transaction Endpoints

### GET /transactions

Get user transactions with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "transactions": [
    {
      "id": "123",
      "userId": "456",
      "type": "TRANSFER",
      "status": "completed",
      "amount": 100.00,
      "currency": "USD",
      "sourceAccountId": "789",
      "destAccountId": "987",
      "fee": 2.50,
      "exchangeRate": 0.855,
      "description": "Transfer to Jane Smith",
      "reference": "REF123456",
      "processedAt": "2023-01-01T00:00:00Z",
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
    }
  ],
  "total": 1
}
```

### GET /transactions/{id}

Get a specific transaction by ID.

**Response:**
```json
{
  "id": "123",
  "userId": "456",
  "type": "TRANSFER",
  "status": "completed",
  "amount": 100.00,
  "currency": "USD",
  "sourceAccountId": "789",
  "destAccountId": "987",
  "fee": 2.50,
  "exchangeRate": 0.855,
  "description": "Transfer to Jane Smith",
  "reference": "REF123456",
  "processedAt": "2023-01-01T00:00:00Z",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

Common HTTP status codes:
- `400`: Bad Request - Invalid input
- `401`: Unauthorized - Missing or invalid authentication
- `403`: Forbidden - Insufficient permissions
- `404`: Not Found - Resource not found
- `429`: Too Many Requests - Rate limit exceeded
- `500`: Internal Server Error - Server error

## Rate Limiting

API endpoints are rate-limited to prevent abuse:
- 100 requests per minute per IP
- 1000 requests per hour per user

Exceeding these limits will result in a 429 Too Many Requests response.

## Webhooks

Globepay provides webhooks for real-time notifications:

### Transfer Status Updates

**Endpoint:** `POST /webhooks/transfer-status`

**Payload:**
```json
{
  "eventType": "transfer.status.updated",
  "transferId": "123",
  "status": "completed",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

### Webhook Security

All webhooks are signed with a secret. Verify the signature by:
1. Extracting the `X-Signature` header
2. Creating a SHA256 HMAC of the payload using your webhook secret
3. Comparing with the signature from the header

## SDKs

Official SDKs are available for:
- JavaScript/Node.js
- Python
- Go
- Java

Check the [SDK repository](https://github.com/your-org/globepay-sdks) for more information.
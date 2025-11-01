# Testing Documentation

This document outlines the testing strategy, frameworks, and procedures for the Globepay application.

## Testing Strategy

Globepay follows a comprehensive testing approach that includes:

1. **Unit Testing** - Testing individual components in isolation
2. **Integration Testing** - Testing interactions between components
3. **End-to-End Testing** - Testing complete user workflows
4. **Performance Testing** - Testing system performance under load
5. **Security Testing** - Testing for vulnerabilities and security issues
6. **Contract Testing** - Testing API contracts and compatibility

## Backend Testing (Go)

### Unit Testing

Unit tests are written using the standard Go testing framework with testify for assertions.

**Running Unit Tests:**
```bash
# Run all unit tests
cd backend
go test ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Test Structure:**
```
backend/
├── internal/
│   ├── api/
│   │   └── handler/
│   │       ├── auth_test.go
│   │       ├── transfer_test.go
│   │       └── user_test.go
│   ├── service/
│   │   ├── auth_service_test.go
│   │   ├── transfer_service_test.go
│   │   └── user_service_test.go
│   └── repository/
│       ├── user_repository_test.go
│       └── transfer_repository_test.go
```

**Example Unit Test:**
```go
func TestTransferService_CreateTransfer(t *testing.T) {
    // Setup
    mockRepo := new(mocks.TransferRepository)
    service := NewTransferService(mockRepo)
    
    // Test data
    transfer := &domain.Transfer{
        UserID: "user123",
        Amount: 100.00,
        // ... other fields
    }
    
    // Mock expectations
    mockRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), transfer).Return(transfer, nil)
    
    // Execute
    result, err := service.CreateTransfer(context.Background(), transfer)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, transfer, result)
    mockRepo.AssertExpectations(t)
}
```

### Integration Testing

Integration tests verify the interaction between different components and external services.

**Running Integration Tests:**
```bash
# Run integration tests
go test -tags=integration ./test/integration/...
```

**Test Structure:**
```
backend/
├── test/
│   ├── integration/
│   │   ├── database_test.go
│   │   ├── redis_test.go
│   │   └── api_test.go
│   └── fixtures/
│       ├── test_data.sql
│       └── mock_responses.json
```

### Test Data Management

Test data is managed through:

1. **Fixtures** - Predefined test data sets
2. **Factories** - Programmatic test data generation
3. **Database Migrations** - Test-specific schema setup

## Frontend Testing (React/TypeScript)

### Unit Testing

Unit tests are written using Jest and React Testing Library.

**Running Unit Tests:**
```bash
# Run all unit tests
cd frontend
npm test

# Run with coverage
npm run test:coverage

# Run specific test file
npm test src/components/auth/Login.test.tsx
```

**Test Structure:**
```
frontend/
├── src/
│   ├── components/
│   │   ├── auth/
│   │   │   ├── Login.test.tsx
│   │   │   └── Signup.test.tsx
│   │   └── transfer/
│   │       └── TransferForm.test.tsx
│   ├── services/
│   │   ├── auth.service.test.ts
│   │   └── transfer.service.test.ts
│   └── store/
│       ├── authSlice.test.ts
│       └── transferSlice.test.ts
```

**Example Unit Test:**
```typescript
import { render, screen, fireEvent } from '@testing-library/react';
import { Provider } from 'react-redux';
import { store } from '../../store';
import Login from './Login';

describe('Login Component', () => {
  test('renders login form', () => {
    render(
      <Provider store={store}>
        <Login />
      </Provider>
    );
    
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument();
  });

  test('shows error message on failed login', async () => {
    render(
      <Provider store={store}>
        <Login />
      </Provider>
    );
    
    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: 'test@example.com' }
    });
    
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: 'wrongpassword' }
    });
    
    fireEvent.click(screen.getByRole('button', { name: /sign in/i }));
    
    expect(await screen.findByText(/invalid credentials/i)).toBeInTheDocument();
  });
});
```

### End-to-End Testing

E2E tests are written using Playwright for browser automation.

**Running E2E Tests:**
```bash
# Run E2E tests
npm run test:e2e

# Run E2E tests in UI mode
npm run test:e2e:ui

# Run specific test
npx playwright test tests/e2e/auth.spec.ts
```

**Test Structure:**
```
frontend/
├── tests/
│   ├── e2e/
│   │   ├── auth.spec.ts
│   │   ├── transfer.spec.ts
│   │   └── dashboard.spec.ts
│   └── fixtures/
│       ├── users.json
│       └── transfers.json
```

**Example E2E Test:**
```typescript
import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test('user can login successfully', async ({ page }) => {
    await page.goto('/login');
    
    await page.fill('input[name="email"]', 'test@example.com');
    await page.fill('input[name="password"]', 'password123');
    await page.click('button[type="submit"]');
    
    await expect(page).toHaveURL('/dashboard');
    await expect(page.getByText('Hello, Test')).toBeVisible();
  });

  test('user cannot login with invalid credentials', async ({ page }) => {
    await page.goto('/login');
    
    await page.fill('input[name="email"]', 'test@example.com');
    await page.fill('input[name="password"]', 'wrongpassword');
    await page.click('button[type="submit"]');
    
    await expect(page.getByText('Invalid credentials')).toBeVisible();
    await expect(page).toHaveURL('/login');
  });
});
```

## Performance Testing

Performance tests are conducted using k6 for load testing.

**Running Performance Tests:**
```bash
# Install k6
npm install -g k6

# Run performance tests
k6 run tests/performance/api-load-test.js
```

**Test Structure:**
```
tests/
├── performance/
│   ├── api-load-test.js
│   ├── transfer-stress-test.js
│   └── user-concurrency-test.js
```

**Example Performance Test:**
```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 }, // Ramp up to 100 users
    { duration: '1m', target: 100 },  // Stay at 100 users
    { duration: '30s', target: 0 },   // Ramp down to 0 users
  ],
};

export default function () {
  const url = 'http://localhost:8080/api/v1/health';
  const res = http.get(url);
  
  check(res, {
    'status is 200': (r) => r.status === 200,
    'response time < 200ms': (r) => r.timings.duration < 200,
  });
  
  sleep(1);
}
```

## Security Testing

Security testing includes static analysis, dependency scanning, and penetration testing.

**Running Security Tests:**
```bash
# Scan for vulnerabilities in dependencies
npm audit
go list -m all | nancy sleuth

# Scan Docker images
trivy image globepay-backend:latest

# Run SAST scan
golangci-lint run --enable-all
npm run lint
```

## Test Environment Setup

### Local Development

For local testing, Docker Compose provides all necessary services:

```bash
# Start test environment
make dev-up

# Run backend tests
make test-backend

# Run frontend tests
make test-frontend
```

### CI/CD Testing

GitHub Actions automatically run tests on every push:

1. **Unit Tests** - Run on all pushes
2. **Integration Tests** - Run on pull requests to main branches
3. **E2E Tests** - Run on staging deployments
4. **Security Scans** - Run weekly and on dependency updates

## Test Data Management

### Test Database

A separate test database is used for integration tests:

```yaml
# docker-compose.test.yml
version: '3.8'
services:
  postgres-test:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: globepay_test
      POSTGRES_USER: testuser
      POSTGRES_PASSWORD: testpass
    ports:
      - "5433:5432"
```

### Test Fixtures

Test fixtures provide consistent test data:

```json
// tests/fixtures/users.json
{
  "validUser": {
    "email": "test@example.com",
    "password": "password123",
    "firstName": "Test",
    "lastName": "User"
  },
  "adminUser": {
    "email": "admin@example.com",
    "password": "admin123",
    "firstName": "Admin",
    "lastName": "User",
    "role": "admin"
  }
}
```

## Code Coverage

Target code coverage thresholds:

- **Unit Tests**: 80% minimum
- **Integration Tests**: 70% minimum
- **E2E Tests**: 60% of critical user flows

**Generating Coverage Reports:**
```bash
# Backend coverage
cd backend
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Frontend coverage
cd frontend
npm run test:coverage
```

## Continuous Testing

### GitHub Actions Workflow

```yaml
# .github/workflows/test.yml
name: Test
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Setup Node
        uses: actions/setup-node@v3
        with:
          node-version: '20'
      - name: Install dependencies
        run: |
          cd backend && go mod download
          cd ../frontend && npm ci
      - name: Run backend tests
        run: cd backend && go test -v ./...
      - name: Run frontend tests
        run: cd frontend && npm test
```

## Test Reporting

Test results are reported through:

1. **Console Output** - Immediate feedback during development
2. **JUnit XML** - Integration with CI/CD systems
3. **HTML Reports** - Detailed test execution reports
4. **Coverage Reports** - Code coverage visualization

## Best Practices

### Backend Testing Best Practices

1. Use table-driven tests for multiple test cases
2. Mock external dependencies
3. Test error cases and edge conditions
4. Use context for cancellation and timeouts
5. Clean up test data after each test

### Frontend Testing Best Practices

1. Test user interactions rather than implementation details
2. Use realistic test data
3. Mock API calls in unit tests
4. Test accessibility
5. Use page objects for E2E tests

### General Testing Best Practices

1. Write tests before code (TDD)
2. Keep tests independent and isolated
3. Use descriptive test names
4. Test both happy paths and error paths
5. Maintain test data consistency
6. Regularly review and update tests

This testing documentation provides a comprehensive guide to testing the Globepay application, ensuring quality and reliability across all components.
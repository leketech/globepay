# Frontend Testing

This directory contains tests for the Globepay frontend application.

## Test Structure

```
tests/
├── unit/                # Unit tests
│   ├── components/      # Component tests
│   ├── services/        # Service tests
│   └── store/           # Redux store tests
├── integration/         # Integration tests
└── e2e/                 # End-to-end tests
```

## Running Tests

### Unit Tests

```bash
# Run all unit tests
npm test

# Run unit tests with coverage
npm run test:coverage

# Run specific test file
npm test src/components/auth/Login.test.tsx
```

### End-to-End Tests

```bash
# Run E2E tests
npm run test:e2e

# Run E2E tests in UI mode
npm run test:e2e:ui
```

## Test Frameworks

- **Jest**: Test runner and assertion library
- **React Testing Library**: For testing React components
- **Playwright**: For end-to-end testing

## Writing Tests

### Unit Tests

Unit tests should focus on testing individual components and functions in isolation.

Example:
```typescript
import { render, screen } from '@testing-library/react';
import Login from '../../src/components/auth/Login';

test('renders login form', () => {
  render(<Login />);
  
  expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
  expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
});
```

### Integration Tests

Integration tests verify that multiple components work together correctly.

### E2E Tests

E2E tests simulate real user interactions with the application.

Example:
```typescript
import { test, expect } from '@playwright/test';

test('user can login successfully', async ({ page }) => {
  await page.goto('/login');
  
  await page.fill('input[name="email"]', 'test@example.com');
  await page.fill('input[name="password"]', 'password123');
  await page.click('button[type="submit"]');
  
  await expect(page).toHaveURL('/dashboard');
});
```

## Test Coverage

Target test coverage:
- Unit tests: 80%+
- Integration tests: 70%+
- E2E tests: 60%+ (critical user flows)

## Mocking

Use mocks for external dependencies:
- API calls
- Browser APIs
- Third-party libraries

## Continuous Integration

Tests are automatically run in CI/CD pipeline:
- On every pull request
- Before deployment to staging
- After deployment to production (smoke tests)
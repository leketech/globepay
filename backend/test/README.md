# Backend Testing

This directory contains all backend tests for the Globepay application.

## Test Structure

```
test/
├── fixtures/           # Test data fixtures
├── integration/        # Integration tests
├── mocks/              # Mock implementations for unit tests
├── unit/               # Unit tests
└── utils/              # Test utilities
```

## Running Tests

### All Tests

```bash
# Run all backend tests
make test-backend

# Run tests with coverage
make test-backend-coverage
```

### Unit Tests

```bash
# Run all unit tests
make test-backend-unit

# Run specific unit test
go test -v ./test/unit/user_service_test.go
```

### Integration Tests

```bash
# Run all integration tests
make test-backend-integration

# Run specific integration test
go test -v ./test/integration/database_test.go
```

## Test Types

### Unit Tests

Unit tests focus on testing individual components in isolation using mocks. They are located in the [unit/](unit/) directory.

### Integration Tests

Integration tests verify that different components work together correctly with real dependencies (database, Redis, etc.). They are located in the [integration/](integration/) directory.

## Fixtures

Test data fixtures are JSON files located in the [fixtures/](fixtures/) directory. These provide consistent test data across different test suites.

## Mocks

Mock implementations are located in the [mocks/](mocks/) directory. These are used in unit tests to isolate the component under test.

## Utilities

Test utilities are located in the [utils/](utils/) directory. These provide helper functions for loading fixtures and other common test operations.
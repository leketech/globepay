# Globepay Backend

This is the backend service for Globepay, a production-ready fintech application for international money transfers.

## Features

- User authentication and management
- Account management
- Money transfer processing
- Transaction history
- Beneficiary management
- Exchange rate services
- Notification services
- Health checks and metrics
- Comprehensive logging

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT
- **API Documentation**: Swagger/OpenAPI
- **Testing**: Go testing framework
- **Containerization**: Docker
- **Orchestration**: Kubernetes

## Getting Started

### Prerequisites

- Go 1.21+
- Docker and Docker Compose
- PostgreSQL (for local development)
- Redis (for local development)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/globepay.git
   cd globepay/backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file based on `.env.example`:
   ```bash
   cp .env.example .env
   # Edit .env with your configurations
   ```

### Running the Application

#### Development Mode

```bash
# Run with hot reload
air

# Or run directly
go run cmd/api/main.go
```

#### Production Mode

```bash
# Build the binary
go build -o bin/api cmd/api/main.go

# Run the binary
./bin/api
```

### Running with Docker

```bash
# Build the Docker image
docker build -t globepay-backend .

# Run the container
docker run -p 8080:8080 globepay-backend
```

### Database Migrations

```bash
# Run migrations up
make migration-up

# Run migrations down
make migration-down

# Create new migration
make migration-create NAME=add_users_table
```

### Testing

```bash
# Run all tests
make test-backend

# Run with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test
go test -v ./internal/service -run TestTransferService
```

## API Documentation

API documentation is available at `/swagger` endpoint when running the backend.

## Project Structure

```
backend/
├── cmd/                # Application entrypoints
│   ├── api/           # Main API server
│   └── migration/     # Database migration tool
├── internal/          # Private application code
│   ├── api/           # HTTP handlers and routes
│   ├── domain/        # Business logic and entities
│   ├── repository/    # Data access layer
│   └── infrastructure/# External dependencies
├── test/              # Test files
└── docs/              # Documentation
```

## Configuration

The application is configured through environment variables. See `.env.example` for all available options.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
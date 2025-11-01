# Quick Start Guide

This guide will help you get Globepay up and running quickly.

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker (24.0+)
- Docker Compose (2.20+)
- Git

## Step 1: Clone the Repository

```bash
git clone https://github.com/your-org/globepay.git
cd globepay
```

## Step 2: Start the Development Environment

```bash
# Make the setup script executable
chmod +x scripts/setup-dev-environment.sh

# Run the setup script
./scripts/setup-dev-environment.sh
```

Alternatively, you can use the Make command:

```bash
make setup
```

## Step 3: Access the Applications

Once the setup is complete, you can access the applications at:

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Grafana**: http://localhost:3001 (admin/admin)
- **Prometheus**: http://localhost:9091

## Step 4: Create Your First User

1. Navigate to the frontend at http://localhost:3000
2. Click "Create Account"
3. Fill in your details and submit
4. You'll be redirected to the dashboard

## Step 5: Make Your First Transfer

1. From the dashboard, click "Send Money"
2. Fill in the recipient details
3. Enter the amount and select currencies
4. Review and confirm the transfer

## Next Steps

- Explore the [API Documentation](http://localhost:8080/swagger)
- Check out the monitoring dashboards in Grafana
- Review the [full documentation](../README.md)
- Learn about [deployment](../DEPLOYMENT.md)

## Troubleshooting

If you encounter any issues:

1. Check the logs: `make dev-logs`
2. Ensure all Docker services are running: `docker-compose ps`
3. Restart the environment: `make dev-down && make dev-up`

For more detailed troubleshooting, see the [Deployment Guide](../DEPLOYMENT.md).
#!/bin/bash

# This script generates mock files for testing using mockery
# Install mockery first: go install github.com/vektra/mockery/v2@latest

echo "Generating mocks for repositories..."
mockery --dir=../internal/repository --name=UserRepoInterface --output=. --outpkg=mocks
mockery --dir=../internal/repository --name=AccountRepoInterface --output=. --outpkg=mocks
mockery --dir=../internal/repository --name=TransferRepoInterface --output=. --outpkg=mocks
mockery --dir=../internal/repository --name=TransactionRepoInterface --output=. --outpkg=mocks

echo "Generating mocks for services..."
mockery --dir=../internal/service --name=UserServiceInterface --output=. --outpkg=mocks
mockery --dir=../internal/service --name=TransferServiceInterface --output=. --outpkg=mocks
mockery --dir=../internal/service --name=TransactionServiceInterface --output=. --outpkg=mocks
mockery --dir=../internal/service --name=AuthServiceInterface --output=. --outpkg=mocks

echo "Mock generation complete!"
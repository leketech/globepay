# Setup Terraform backend resources (S3 bucket and DynamoDB table)

Write-Host "Setting up Terraform backend resources..."

# Create S3 bucket for state storage
Write-Host "Creating S3 bucket for Terraform state..."
aws s3api create-bucket `
  --bucket globepay-terraform-state-prod `
  --region us-east-1

# Enable versioning on the S3 bucket
Write-Host "Enabling versioning on S3 bucket..."
aws s3api put-bucket-versioning `
  --bucket globepay-terraform-state-prod `
  --versioning-configuration Status=Enabled

# Create DynamoDB table for state locking
Write-Host "Creating DynamoDB table for state locking..."
aws dynamodb create-table `
  --table-name globepay-terraform-locks `
  --attribute-definitions AttributeName=LockID,AttributeType=S `
  --key-schema AttributeName=LockID,KeyType=HASH `
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 `
  --region us-east-1

Write-Host "Terraform backend setup complete!"
Write-Host "S3 bucket: globepay-terraform-state-prod"
Write-Host "DynamoDB table: globepay-terraform-locks"
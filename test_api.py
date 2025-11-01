import requests
import json
import time

# Register a new user with unique email
timestamp = int(time.time())
register_data = {
    "email": f"test{timestamp}@example.com",
    "password": "password123",
    "firstName": "Test",
    "lastName": "User"
}

print("Registering a new user...")
response = requests.post(
    "http://localhost:8080/api/v1/auth/register",
    headers={"Content-Type": "application/json"},
    json=register_data
)

print(f"Status code: {response.status_code}")
print(f"Response: {response.text}")

if response.status_code in [200, 201]:
    response_data = response.json()
    token = response_data.get("token")
    print(f"Token: {token}")
    
    # Get transfers
    print("Getting transfers...")
    transfers_response = requests.get(
        "http://localhost:8080/api/v1/transfers",
        headers={"Authorization": f"Bearer {token}"}
    )
    
    print(f"Transfers status code: {transfers_response.status_code}")
    print(f"Transfers response: {transfers_response.text}")
else:
    print("Failed to register user")
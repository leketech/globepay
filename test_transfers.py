import requests
import json

# Test login
login_data = {
    "email": "test@example.com",
    "password": "password123"
}

print("Logging in...")
response = requests.post(
    "http://localhost:8080/api/v1/auth/login",
    headers={"Content-Type": "application/json"},
    json=login_data
)

print(f"Status code: {response.status_code}")
print(f"Response: {response.text}")

if response.status_code == 200:
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
    print("Failed to login")
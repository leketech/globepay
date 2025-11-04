import requests
import json

url = "http://localhost:8080/api/v1/auth/register"
headers = {
    "Content-Type": "application/json"
}
data = {
    "email": "test@example.com",
    "password": "password123",
    "firstName": "Test",
    "lastName": "User"
}

try:
    response = requests.post(url, headers=headers, data=json.dumps(data))
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")
except Exception as e:
    print(f"Error: {e}")
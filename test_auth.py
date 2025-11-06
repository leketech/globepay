import requests
import json

# Test signup
signup_url = "http://localhost:8080/api/v1/auth/register"
signup_data = {
    "email": "testuser@example.com",
    "password": "securepassword123",
    "firstName": "Test",
    "lastName": "User",
    "phoneNumber": "+1234567890",
    "dateOfBirth": "1990-01-01",
    "country": "USA"
}

print("Testing signup...")
try:
    response = requests.post(signup_url, json=signup_data)
    print(f"Signup response status: {response.status_code}")
    print(f"Signup response: {response.text}")
    
    if response.status_code == 201:
        print("Signup successful!")
        signup_response = response.json()
        token = signup_response.get("token")
        
        # Test login
        login_url = "http://localhost:8080/api/v1/auth/login"
        login_data = {
            "email": "testuser@example.com",
            "password": "securepassword123"
        }
        
        print("\nTesting login...")
        response = requests.post(login_url, json=login_data)
        print(f"Login response status: {response.status_code}")
        print(f"Login response: {response.text}")
        
        if response.status_code == 200:
            print("Login successful!")
        else:
            print("Login failed!")
    else:
        print("Signup failed!")
        
except Exception as e:
    print(f"Error: {e}")
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	url := "http://localhost:8080/api/v1/auth/register"
	
	data := RegisterRequest{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "Test",
		LastName:  "User",
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	
	fmt.Printf("Sending JSON: %s\n", string(jsonData))
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	
	// Read response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Printf("Response: %s\n", buf.String())
}
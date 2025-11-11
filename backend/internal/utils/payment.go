package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GeneratePaymentLink generates a unique payment link for a money request
func GeneratePaymentLink(requestID string) string {
	// Generate a random token
	token := generateRandomToken(32)

	// Create the payment link
	return fmt.Sprintf("/pay/%s/%s", requestID, token)
}

// generateRandomToken generates a random token of specified length
func generateRandomToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to UUID if random generation fails
		return GenerateUUID()
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

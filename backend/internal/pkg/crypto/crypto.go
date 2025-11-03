package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io"
)

// Crypto provides cryptographic functions
type Crypto struct {
	encryptionKey []byte
}

// NewCrypto creates a new Crypto instance
func NewCrypto(secretKey string) *Crypto {
	// Generate a 32-byte key from the secret key using SHA-256
	key := sha256.Sum256([]byte(secretKey))
	return &Crypto{
		encryptionKey: key[:],
	}
}

// Encrypt encrypts plaintext using AES-GCM
func (c *Crypto) Encrypt(plaintext string) (string, error) {
	// Convert string to []byte
	plaintextBytes := []byte(plaintext)
	
	// Create a new AES cipher
	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES-GCM
func (c *Crypto) Decrypt(ciphertext string) (string, error) {
	// Decode from base64
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Get the nonce size
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %w", err)
	}

	return string(plaintext), nil
}

// HashSHA256 generates a SHA-256 hash of the input
func (c *Crypto) HashSHA256(input []byte) []byte {
	hash := sha256.Sum256(input)
	return hash[:]
}

// GenerateRandomBytes generates random bytes
func (c *Crypto) GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return b, nil
}

// GenerateRandomString generates a random string of specified length
func (c *Crypto) GenerateRandomString(length int) (string, error) {
	bytes, err := c.GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// ConstantTimeCompare compares two byte slices in constant time
func (c *Crypto) ConstantTimeCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	return subtle.ConstantTimeCompare(a, b) == 1
}

// MaskString masks a string for logging purposes
func (c *Crypto) MaskString(input string, visibleChars int) string {
	if len(input) <= visibleChars*2 {
		return "***"
	}
	
	start := input[:visibleChars]
	end := input[len(input)-visibleChars:]
	return start + "***" + end
}
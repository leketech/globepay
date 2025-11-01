package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware handles JWT authentication
func AuthMiddleware(secret string) gin.HandlerFunc {
	fmt.Printf("AuthMiddleware created with secret: %s\n", secret)
	
	return func(c *gin.Context) {
		fmt.Printf("AuthMiddleware called with secret: %s\n", secret)
		
		authHeader := c.GetHeader("Authorization")
		fmt.Printf("Authorization header: %s\n", authHeader)
		
		if authHeader == "" {
			fmt.Println("Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
				"code":  "MISSING_AUTH_HEADER",
			})
			c.Abort()
			return
		}

		// Check if the header has the correct format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("Authorization header format is invalid")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
				"code":  "INVALID_AUTH_HEADER",
			})
			c.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Printf("Token string: %s\n", tokenString)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			fmt.Printf("Token parsing error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Printf("Claims: %+v\n", claims)
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
			fmt.Printf("Set user_id in context: %v\n", claims["user_id"])
		} else {
			fmt.Println("Failed to extract claims")
		}

		c.Next()
	}
}
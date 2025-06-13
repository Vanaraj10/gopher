package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a Gin middleware function that validates JWT tokens from the "Authorization" header.
// It expects the header in the format "Bearer <token>". If the token is missing, invalid, or expired,
// the middleware aborts the request with a 401 Unauthorized response. On successful validation, it extracts
// the "user_id" claim from the token, stores it in the Gin context as "userID", and allows the request to proceed.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			// If the token is invalid or expired, return an unauthorized error
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "Invalid or expired token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		userID := int(claims["user_id"].(float64)) // Convert float64 to int
		c.Set("userID", userID) // Store user ID in context for later use
		c.Next() // Proceed to the next handler
	}
}
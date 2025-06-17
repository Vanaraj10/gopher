package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
// Package utils provides utility functions for generating JWT tokens for email verification and user authentication.
func GenerateEmailToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"exp"  :time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// GenerateAuthToken creates a JWT token for user authentication.
func GenerateAuthToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

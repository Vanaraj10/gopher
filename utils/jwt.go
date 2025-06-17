package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateEmailToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"exp"  :time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
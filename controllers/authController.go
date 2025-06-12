package controllers

import (
	"net/http"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",input.Username, input.Email, string(hashedPassword))
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {

			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
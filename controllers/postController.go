package controllers

import (
	"fmt"
	"net/http"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetInt("user_id")
	fmt.Println(userID) // Retrieve user ID from context

	_, err := config.DB.Exec("INSERT INTO posts (user_id, content) VALUES ($1, $2)", userID, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}
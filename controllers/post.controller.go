package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Vanaraj10/gopher-backend/config"
	"github.com/Vanaraj10/gopher-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	var postCollection = config.GetCollection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	author := user.(models.User)
	post.ID = primitive.NewObjectID()
	post.AuthorID = author.ID
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := postCollection.InsertOne(ctx, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

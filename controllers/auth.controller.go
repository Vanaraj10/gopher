package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Vanaraj10/gopher-backend/config"
	"github.com/Vanaraj10/gopher-backend/models"
	"github.com/Vanaraj10/gopher-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func Signup(c *gin.Context) {
	var userCollection = config.GetCollection("users")
	var user models.User

	// Bind the JSON input to the user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}
	// Validate the user model
	if err := validate.Struct(user); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	// Check if the user already exists by email
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user with the hashed password and other details
	user.ID = primitive.NewObjectID()
	user.Password = hashedPassword
	user.Verified = false
	user.CreatedAt = time.Now()

	// Insert the new user into the database
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"+err.Error()})
		return
	}

	token, err := utils.GenerateEmailToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Token creation failed",
		})
		return
	}
	if err := utils.SendVerificationEmail(user.Email,token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
        return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func VerifyEmail(c *gin.Context) {
	var userCollection = config.GetCollection("users")
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
        return
	}

	token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC) ; !ok {
			 return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")),nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["email"]==nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"Invalid token data",
		})
		return
	}
	email := claims["email"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	filter := bson.M{"email":email}
	update := bson.M{"$set":bson.M{"verified":true}}

	res, err := userCollection.UpdateOne(ctx,filter,update)
	if err != nil || res.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user"})
        return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Email verified successfully",
	})
}
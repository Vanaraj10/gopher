package main

import (
	"log"
	"os"

	"github.com/Vanaraj10/gopher-backend/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load() ; err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	router := gin.Default()

	router.GET("/hello",func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	port := os.Getenv("PORT")
	router.Run(":" + port) // Run the server on the specified port
}
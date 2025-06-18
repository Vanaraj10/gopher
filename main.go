package main

import (
	"log"
	"os"

	"github.com/Vanaraj10/gopher-backend/config"
	"github.com/Vanaraj10/gopher-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load() ; err != nil {
		log.Fatal("Error loading .env file")
	}
	
	config.ConnectDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	
	router.Run(":" + port) // Run the server on the specified port
}
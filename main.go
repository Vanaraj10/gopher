package main

import (
	"log"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/Vanaraj10/social-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.ProtectedRoutes(r)
	routes.PostRoutes(r)
	r.Run(":8080") // Start the server on port 8080
	log.Println("Server running on port 8080")
}
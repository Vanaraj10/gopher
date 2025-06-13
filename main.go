package main

import (
	"log"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/Vanaraj10/social-backend/middleware"
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
	routes.FollowRoutes(r)
	routes.RegisterCommentRoutes(r, middleware.AuthMiddleware())
	routes.RegisterUserRoutes(r, middleware.AuthMiddleware())
	r.Run(":8080") // Start the server on port 8080
	log.Println("Server running on port 8080")
}
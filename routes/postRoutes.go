package routes

import (
	"github.com/Vanaraj10/social-backend/controllers"
	"github.com/Vanaraj10/social-backend/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleware()) // Apply authentication middleware to all post routes
	posts.POST("/create", controllers.CreatePost) // Create a new post
}
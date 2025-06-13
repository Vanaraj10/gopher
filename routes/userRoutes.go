package routes

import (
	"github.com/Vanaraj10/social-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, authMiddleware gin.HandlerFunc) {
	// Register the user routes under the "/users" group
	users := r.Group("/users")
	users.Use(authMiddleware) // Apply authentication middleware to all user routes
	{
		users.GET("/:username", controllers.GetUserProfile) // Get user profile by username
	}
}
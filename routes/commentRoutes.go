package routes

import (
	"github.com/Vanaraj10/social-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(r *gin.Engine,authMiddleware gin.HandlerFunc) {
	// Register the comment routes under the "/comments" group
	comments := r.Group("/comments")
	comments.Use(authMiddleware) // Apply authentication middleware to all comment routes
	{
		comments.POST("/:post_id", controllers.AddComment) // Create a new comment
		comments.GET("/:post_id", controllers.GetCommentsByPost)    // Get comments for a specific post   
	}
}
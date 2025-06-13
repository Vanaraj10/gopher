package routes

import (
	"github.com/Vanaraj10/social-backend/controllers"
	"github.com/Vanaraj10/social-backend/middleware"
	"github.com/gin-gonic/gin"
)

func FollowRoutes(r *gin.Engine){
	follow := r.Group("/follow")
	
	// Apply authentication middleware to all follow routes
	follow.Use(middleware.AuthMiddleware())
	
	follow.POST("/:id", controllers.FollowUser) // Follow a user
	follow.DELETE("/:id", controllers.UnfollowUser) // Unfollow a user

	feed := r.Group("/feed")
	feed.Use(middleware.AuthMiddleware())
	feed.GET("/", controllers.GetFeed) // Get the feed of posts from followed users
}
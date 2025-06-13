package routes

import (
	"github.com/Vanaraj10/social-backend/middleware"
	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me"	,func ( c *gin.Context)  {
		userID := c.MustGet("userID").(int) // Retrieve user ID from context
		c.JSON(200,gin.H{
			"user_id": userID,
			"message": "This is a protected route",
		})
	})
}
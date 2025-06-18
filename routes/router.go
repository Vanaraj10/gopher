package routes

import (
	"github.com/Vanaraj10/gopher-backend/controllers"
	"github.com/Vanaraj10/gopher-backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/signup",controllers.Signup)
		api.GET("/verify",controllers.VerifyEmail)
		api.POST("/login", controllers.Login)
		api.GET("/me",middleware.AuthMiddleware(), controllers.Me)
		api.POST("/posts", middleware.AuthMiddleware(), controllers.CreatePost)
	}
}
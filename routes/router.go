package routes

import (
	"github.com/Vanaraj10/gopher-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/signup",controllers.Signup)
	}
}
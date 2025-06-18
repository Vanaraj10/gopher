// RegisterRoutes sets up the API routes for the application using the provided gin.Engine.
// It registers the following endpoints under the "/api" group:
//   - POST   /signup    : Handles user registration via controllers.Signup
//   - GET    /verify    : Handles email verification via controllers.VerifyEmail
//   - POST   /login     : Handles user login via controllers.Login
//   - GET    /me        : Returns current user info, requires authentication via middleware.AuthMiddleware
//   - POST   /posts     : Creates a new post, requires authentication via middleware.AuthMiddleware
//   - GET    /posts     : Retrieves all posts via controllers.GetAllPosts
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
		api.GET("/posts",controllers.GetAllPosts)
		api.GET("/posts/:id", controllers.GetPostByID)
		api.PUT("/posts/:id", middleware.AuthMiddleware(), controllers.UpdatePost)
	}
}
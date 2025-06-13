package controllers

import (
	"net/http"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/Vanaraj10/social-backend/models"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	// Logic to retrieve user profile by username

	var user models.User
	err := config.DB.QueryRow(
		`SELECT id, username, bio, email, created_at FROM users WHERE username = $1`,username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Bio,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	 // Get Post Count
    var postCount int
    err = config.DB.QueryRow(`
        SELECT COUNT(*) FROM posts WHERE user_id = $1
    `, user.ID).Scan(&postCount)

    // Followers count
    var followers int
    err = config.DB.QueryRow(`
        SELECT COUNT(*) FROM follows WHERE following_id = $1
    `, user.ID).Scan(&followers)

    // Following count
    var following int
    err = config.DB.QueryRow(`
        SELECT COUNT(*) FROM follows WHERE follower_id = $1
    `, user.ID).Scan(&following)

	c.JSON(http.StatusOK,gin.H{
		"id": user.ID,
		"username": user.Username,
		"bio": user.Bio,
		"email": user.Email,
		"created_at": user.CreatedAt,
		"post_count": postCount,
		"followers": followers,
		"following": following,
		"message": "User profile retrieved successfully",
	})
}
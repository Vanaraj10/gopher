package controllers

import (
	"net/http"
	"strconv"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/Vanaraj10/social-backend/models"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	// Logic to add a comment to a post
	// This function will handle the request to add a comment
	// It should validate the input, check if the user is authenticated,
	// and then save the comment to the database.
	userID := c.GetInt("user_id")                   // Retrieve user ID from context
	postID, err := strconv.Atoi(c.Param("post_id")) // Get post ID from URL parameters

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	_, err = config.DB.Exec(
		`INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3)`,
		postID, userID, comment.Content,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

func GetCommentsByPost(c *gin.Context) {
	// Logic to retrieve comments for a specific post
	// This function will handle the request to get comments for a post
	// It should query the database and return the comments in JSON format.
	postID, err := strconv.Atoi(c.Param("post_id")) // Get post ID from URL parameters
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}
	rows, err := config.DB.Query(
		`
		SELECT id, post_id, user_id, content, created_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC
		`, postID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	var comments []models.Comment
	for rows.Next() {
		var cmt models.Comment
		if err := rows.Scan(&cmt.ID, &cmt.PostID, &cmt.UserID, &cmt.Content, &cmt.CreatedAt); err == nil {
			comments = append(comments, cmt)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan comment"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func DeleteComment(c *gin.Context) {
	// Logic to delete a comment
	// This function will handle the request to delete a comment
	// It should check if the user is authorized to delete the comment
	// and then remove it from the database.
	userID := c.GetInt("user_id")                 // Retrieve user ID from context
	commentID, err := strconv.Atoi(c.Param("id")) // Get comment ID from URL parameters
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}
	result, err := config.DB.Exec(
		`DDELETE FROM comments
		 WHERE id = $1 AND user_id = $2`,
		commentID, userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Comment not found or not authorized to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Vanaraj10/social-backend/config"
    "github.com/Vanaraj10/social-backend/models"
)

func GetFeed(c *gin.Context) {
    userID := c.GetInt("user_id")

    rows, err := config.DB.Query(`
        SELECT posts.id, posts.user_id, posts.content, posts.created_at
        FROM posts
        JOIN follows ON posts.user_id = follows.followed_id
        WHERE follows.follower_id = $1
        ORDER BY posts.created_at DESC
    `, userID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feed"})
        return
    }
    defer rows.Close()

    var feed []models.Post
    for rows.Next() {
        var post models.Post
        if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt); err == nil {
            feed = append(feed, post)
        }
    }

    c.JSON(http.StatusOK, gin.H{"feed": feed})
}

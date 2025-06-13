package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Vanaraj10/social-backend/config"
)

func FollowUser(c *gin.Context) {
    followerID := c.GetInt("user_id")
    followedID, err := strconv.Atoi(c.Param("id"))
    if err != nil || followerID == followedID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follow request"})
        return
    }

    _, err = config.DB.Exec(
        "INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
        followerID, followedID,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Now following user"})
}

func UnfollowUser(c *gin.Context) {
    followerID := c.GetInt("user_id")
    followedID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unfollow request"})
        return
    }

    _, err = config.DB.Exec(
        "DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2",
        followerID, followedID,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Unfollowed user"})
}

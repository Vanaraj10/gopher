package main

import (
	"log"

	"github.com/Vanaraj10/social-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()
	
	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"Message":"Social API is running",
		})
	})

	if err := r.Run(":8080") ; err != nil{
		log.Fatal("Failed to run server:",err)
	}
}
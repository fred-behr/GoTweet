package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Simple GET endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.Run(":8080")
}

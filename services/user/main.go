package main

import (
	"GoTweet/services/user/database"
	"GoTweet/services/user/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Simple GET endpoint
	dbClient, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	r.GET("/readiness", func(c *gin.Context) {
		ready := dbClient.Ready()
		if ready {
			c.JSON(200, models.Health{Status: "OK"})
			return
		}
		c.JSON(500, models.Health{Status: "Failure"})
	})

	r.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, models.Health{Status: "OK"})
	})

	r.Run(":8080")
}

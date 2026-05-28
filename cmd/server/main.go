package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anandudevops/aegis/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var startTime = time.Now()

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables")
	}

	

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Health
	r.GET("/health", func(c *gin.Context) {
		response.Success(c, 200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
			"uptime":  fmt.Sprintf("%.0fs", time.Since(startTime).Seconds()),
		})
	})
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server: %v", err)
	}
}

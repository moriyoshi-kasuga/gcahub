package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CORS := os.Getenv("CORS")
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{CORS}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the API server.")
	})

	r.Run()
}

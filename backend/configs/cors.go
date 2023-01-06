package config

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://bard.pages.dev"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}

	log.Default().Println("CORS Config: ", config)

	return config
}

package main

import (
	"log"
	"os"
	"ybigta/bard-backend/api"
	config "ybigta/bard-backend/configs"
	db_repository "ybigta/bard-backend/repository/db"
	session_repository "ybigta/bard-backend/repository/session"
	util "ybigta/bard-backend/utils"

	"github.com/joho/godotenv"
)

func loadEnv() {
	// Load env variables

	gin_mode := os.Getenv("GIN_MODE")

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// debug mode
	if gin_mode == "" {
		err = godotenv.Load(".env.development")
		if err != nil {
			panic("Error loading .env.development file")
		}
	}
	if gin_mode == "release" {
		err = godotenv.Load(".env.production")
		if err != nil {
			panic("Error loading .env.production file")
		}
	}
}

func main() {

	loadEnv()

	dsn := config.GetDSN()
	db := db_repository.NewDatabase(dsn)
	db.Connect()

	session_repository.ConnectSessionStore()

	util.InitializeBLIP()
	util.InitializeDALLE()
	config.LoadGPT3Config()

	log.Default().Println("GPT3 Config Loaded")

	util.InitializeGPT3()

	log.Default().Println("GPT3 Initialized")

	util.InitializeS3Client()

	log.Default().Println("S3 Client Initialized")

	router := api.InitRouter()

	log.Default().Println("Router Initialized")

	err := router.Run()

	log.Default().Println("Server running on port 8080")

	if err != nil {
		log.Fatal(err)
		panic("failed to run server")
	}

}

package config

import (
	"os"
)

func GetDSN() string {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	db_url := username + ":" + password + "@(" + db_host + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"

	return db_url
}

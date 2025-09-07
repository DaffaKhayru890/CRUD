package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DaffaKhayru890/src/configs"
	"github.com/DaffaKhayru890/src/models"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1]

	envPath, err := filepath.Abs("../.env") // migrations/ -> ../.env
	if err != nil {
		log.Fatal("Failed to get absolute path for .env:", err)
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if strings.ToLower(args) == "up" {
		err := configs.DB().AutoMigrate(&models.Contacts{})

		if err != nil {
			log.Fatal("Can not migrate database")
		}
	}else if strings.ToLower(args) == "down" {
		err := configs.DB().Migrator().DropTable(&models.Contacts{})

		if err != nil {
			log.Fatal("Can not migrate database")
		}
	}
}
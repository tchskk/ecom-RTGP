package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/tchskk/ecom-RTGP/backend/config"
	"github.com/tchskk/ecom-RTGP/backend/models"
	"github.com/tchskk/ecom-RTGP/backend/routes"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	// Auto-migrate
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tchskk/ecom-RTGP/backend/db"
	"github.com/tchskk/ecom-RTGP/backend/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize database tables
	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Setup router
	router := routes.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

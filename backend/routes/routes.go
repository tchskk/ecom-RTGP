package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tchskk/ecom-RTGP/backend/handlers"
	"github.com/tchskk/ecom-RTGP/backend/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Vite default port
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Public routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", handlers.GetProfile)
	}

	// Admin routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		admin.GET("/users", handlers.GetAllUsers)
	}

	return r
}

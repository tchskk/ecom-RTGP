package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/tchskk/ecom-RTGP/backend/controllers"
	"github.com/tchskk/ecom-RTGP/backend/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware("admin"))
		{
			admin.GET("/users", controllers.GetAllUsers)
		}
	}

	return r
}
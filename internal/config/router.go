package config

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hnex.com/internal/handlers"
)

func SetupRouter(app *gin.Engine) *gin.Engine {
	env := LoadEnv()
	db, err := ConnectDB(env.DATABASE_URL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	api := app.Group("/api")
	{
		authHandler := handlers.AuthHandler{Db: db}
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/logout", authHandler.Logout)
			auth.POST("/refresh", authHandler.RefreshToken)
		}
	}

	return app
}

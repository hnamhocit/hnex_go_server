package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hnex.com/internal/config"
	"hnex.com/internal/handlers"
	"hnex.com/internal/middlewares"
	"hnex.com/internal/repositories"
	"hnex.com/internal/services"
)

func Start(env *config.Env, db *gorm.DB, hostname string) {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	api := app.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})

		userRepository := repositories.UserRepository{DB: db}
		authRepository := repositories.AuthRepository{DB: db}

		authService := services.AuthService{
			Repository:     authRepository,
			UserRepository: userRepository,
		}

		authHandler := handlers.AuthHandler{Service: authService}
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/logout", middlewares.AccessTokenMiddleware, authHandler.Logout)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		userService := services.UserService{
			Repository: userRepository,
		}

		userHandler := handlers.UserHandler{Service: userService}
		user := api.Group("/user")
		{
			user.GET("/profile", middlewares.AccessTokenMiddleware, userHandler.GetProfile)
			user.GET("/:id", userHandler.GetUser)
		}
	}

	log.Printf("Starting server on port %s:%d", hostname, env.PORT)
	app.Run(fmt.Sprintf(":%d", env.PORT))
}

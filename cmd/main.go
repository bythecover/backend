package main

import (
	poll_handler "bythecover/backend/internal/adapters/handler/poll"
	user_handler "bythecover/backend/internal/adapters/handler/user"
	poll_repository "bythecover/backend/internal/adapters/repository/poll"
	"bythecover/backend/internal/adapters/repository/postgres"
	user_repository "bythecover/backend/internal/adapters/repository/user"
	"bythecover/backend/internal/core/services"

	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        // c.Header("Access-Control-Allow-Credentials", "true")
        // c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        // c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
	}
}

func main() {
	dbConnection := postgres.NewPostgresDatabase()

	userRepo := user_repository.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userHandler := user_handler.NewUserHttpHandler(userService)

	pollRepo := poll_repository.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo)
	pollHandler := poll_handler.NewPollHttpHandler(pollService)

	route := gin.Default()
	route.Use(SetupCors())

	userHandler.RegisterRoutes(route)
	pollHandler.RegisterRoutes(route)
	
	route.Run(":8080")
}
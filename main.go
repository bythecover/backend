package main

import (
	htmx_handler "bythecover/backend/internal/adapters/handler/htmx"
	poll_handler "bythecover/backend/internal/adapters/handler/poll"
	user_handler "bythecover/backend/internal/adapters/handler/user"
	poll_repository "bythecover/backend/internal/adapters/repository/poll"
	"bythecover/backend/internal/adapters/repository/postgres"
	user_repository "bythecover/backend/internal/adapters/repository/user"
	"bythecover/backend/internal/core/services"
	"bythecover/backend/internal/core/services/htmx"

	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "null")
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

	htmxService := htmx.NewHtmxService()
	htmxHandler := htmx_handler.NewHtmxHttpHandler(htmxService, pollService)

	route := gin.Default()
	route.Use(SetupCors())

	userHandler.RegisterRoutes(route)
	pollHandler.RegisterRoutes(route)
	htmxHandler.RegisterRoutes(route)

	route.Run(":8080")
}

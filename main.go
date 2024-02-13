package main

import (
	"bythecover/backend/internal/adapters/handler"
	"bythecover/backend/internal/adapters/persistence"
	"bythecover/backend/internal/core/services"

	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "null")
	}
}

func main() {
	dbConnection := persistence.NewPostgresDatabase()

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHttpHandler(userService)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo)
	pollHandler := handler.NewPollHttpHandler(pollService)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)
	voteService := services.NewVoteService(voteRepo)

	htmxService := services.NewHtmxService(voteService)
	htmxHandler := handler.NewHtmxHttpHandler(htmxService, pollService)

	route := gin.Default()
	route.Use(SetupCors())

	userHandler.RegisterRoutes(route)
	pollHandler.RegisterRoutes(route)
	htmxHandler.RegisterRoutes(route)

	route.Run(":8080")
}

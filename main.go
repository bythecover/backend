package main

import (
	htmx_handler "bythecover/backend/internal/adapters/handler/htmx"
	poll_handler "bythecover/backend/internal/adapters/handler/poll"
	user_handler "bythecover/backend/internal/adapters/handler/user"
	poll_repository "bythecover/backend/internal/adapters/repository/poll"
	"bythecover/backend/internal/adapters/repository/postgres"
	user_repository "bythecover/backend/internal/adapters/repository/user"
	vote_repository "bythecover/backend/internal/adapters/repository/vote"
	"bythecover/backend/internal/core/services"
	"bythecover/backend/internal/core/services/htmx"

	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "null")
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

	voteRepo := vote_repository.NewPollPostgresRepository(dbConnection)
	voteService := services.NewVoteService(voteRepo)

	htmxService := htmx.NewHtmxService(voteService)
	htmxHandler := htmx_handler.NewHtmxHttpHandler(htmxService, pollService)

	route := gin.Default()
	route.Use(SetupCors())

	userHandler.RegisterRoutes(route)
	pollHandler.RegisterRoutes(route)
	htmxHandler.RegisterRoutes(route)

	route.Run(":8080")
}

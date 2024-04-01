package main

import (
	http_adapter "bythecover/backend/internal/adapters/http"
	"bythecover/backend/internal/adapters/persistence"
	"bythecover/backend/internal/core/services"
	"log"
	"net/http"
)

func main() {
	dbConnection := persistence.NewPostgresDatabase()

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userHandler := http_adapter.NewUserHttpHandler(userService)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo)
	pollHandler := http_adapter.NewPollHttpHandler(pollService)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)
	voteService := services.NewVoteService(voteRepo)

	htmxService := services.NewHtmxService(voteService)
	htmxHandler := http_adapter.NewHtmxHttpHandler(htmxService, pollService)

	router := http.NewServeMux()

	userHandler.RegisterRoutes(router)
	pollHandler.RegisterRoutes(router)
	htmxHandler.RegisterRoutes(router)

	server := http.Server{
		Handler: http_adapter.AllowCors(http_adapter.Logger(router)),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

package main

import (
	adapters "bythecover/backend/internal/adapters/http"
	"bythecover/backend/internal/adapters/persistence"
	"bythecover/backend/internal/core/services"
	"bythecover/backend/internal/core/services/authenticator"
	"bythecover/backend/internal/core/services/sessions"
	"log"
	"net/http"

	"github.com/goloop/env"
)

func main() {
	if err := env.Update(".env"); err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()

	dbConnection := persistence.NewPostgresDatabase()
	sessionStore := make(sessions.MemoryStore)
	sessions.CreateStore(sessionStore)

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userAdapter := adapters.NewUserHttpAdapter(userService)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo, voteRepo)
	adapters.NewPollHttpAdapter(pollService, router)

	userAdapter.RegisterRoutes(router)

	sessionHandler := adapters.HandlerWithSession(sessionStore)
	middlewareStack := adapters.CreateStack(sessionHandler, adapters.AllowCors, adapters.Logger)

	authService, _ := authenticator.New()

	adapters.NewLoginHttpAdapter(authService, router)

	callbackAdapter := adapters.NewCallbackHttpAdapter(authService)
	callbackAdapter.RegisterRoutes(router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

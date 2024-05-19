package main

import (
	adapters "bythecover/backend/internal/adapters/http"
	"bythecover/backend/internal/adapters/persistence"
	"bythecover/backend/internal/core/services"
	"log"
	"net/http"

	"github.com/goloop/env"
)

func main() {
	if err := env.Update(".env"); err != nil {
		log.Fatalln(err)
	}

	dbConnection := persistence.NewPostgresDatabase()
	sessionStore := make(adapters.SessionStore)

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userAdapter := adapters.NewUserHttpAdapter(userService)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo, voteRepo)
	pollAdapter := adapters.NewPollHttpAdapter(pollService)

	router := http.NewServeMux()

	userAdapter.RegisterRoutes(router)
	pollAdapter.RegisterRoutes(router)

	sessionHandler := adapters.HandlerWithSession(sessionStore)
	middlewareStack := adapters.CreateStack(sessionHandler, adapters.AllowCors, adapters.Logger)

	// TODO: fix this
	auth, _ := services.New()

	loginAdapter := adapters.NewLoginHttpAdapter(auth)
	loginAdapter.RegisterRoutes(router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

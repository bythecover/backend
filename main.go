package main

import (
	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/http/routers"
	"github.com/bythecover/backend/persistence"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
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

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo, voteRepo)

	sessionHandler := middleware.HandlerWithSession(sessionStore)
	middlewareStack := middleware.CreateStack(sessionHandler, middleware.AllowCors, middleware.Logger)

	authService, _ := authenticator.New()

	routers.NewLoginHttpAdapter(authService, router)
	routers.NewCallbackHttpAdapter(authService, router)
	routers.NewPollHttpAdapter(pollService, router)
	routers.NewUserHttpAdapter(userService, router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

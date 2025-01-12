package main

import (
	"net/http"
	"os"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/persistence"
	"github.com/bythecover/backend/routers"
	"github.com/bythecover/backend/sessions"
	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/goloop/env"
)

func main() {
	isProd := os.Getenv("ENVIRONMENT") == "PROD"

	if !isProd {
		if err := env.Update(".env"); err != nil {
			logger.Error.Fatalln(err)
		}
	}

	cld, err := cloudinary.New()
	if err != nil {
		logger.Error.Fatalln("Unable to instantiate cloudinary: ", err.Error())
	}

	// I/O
	router := http.NewServeMux()
	dbConnection := persistence.NewPostgresDatabase()

	// User Session/Auth
	sessionStore := sessions.MemoryStore{}
	sessions.CreateStore(sessionStore)
	authService, err := authenticator.New()
	if err != nil {
		logger.Error.Fatalln("Failed to create auth service: ", err)
	}

	// Repos
	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	voteRepo := persistence.NewVotePostgresRepository(dbConnection)
	pollRepo := persistence.NewPollPostgresRepository(dbConnection)

	// Middleware
	sessionHandler := routers.HandlerWithSession(sessionStore)
	middlewareStack := routers.CreateStack(sessionHandler, routers.AllowCors, routers.Logger)

	// Setting up endpoints
	routers.NewAuthorHttpAdapter(router, pollRepo)
	routers.NewLoginHttpAdapter(authService, router)
	routers.NewCallbackHttpAdapter(authService, userRepo, router)
	routers.NewPollHttpAdapter(pollRepo, voteRepo, cld, router)
	routers.NewUserHttpAdapter(userRepo, router)
	routers.NewStaticHttpAdapter(router)
	routers.Home(router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	logger.Error.Fatal(server.ListenAndServe())
}

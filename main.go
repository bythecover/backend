package main

import (
	"net/http"
	"os"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/http/routers"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/persistence"
	"github.com/bythecover/backend/services"
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

	router := http.NewServeMux()

	dbConnection := persistence.NewPostgresDatabase()

	sessionStore := sessions.MemoryStore{}
	sessions.CreateStore(sessionStore)

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)
	pollRepo := persistence.NewPollPostgresRepository(dbConnection)

	sessionHandler := middleware.HandlerWithSession(sessionStore)
	middlewareStack := middleware.CreateStack(sessionHandler, middleware.AllowCors, middleware.Logger)

	authService, _ := authenticator.New()

	routers.NewAuthorHttpAdapter(router, pollRepo)
	routers.NewLoginHttpAdapter(authService, router)
	routers.NewCallbackHttpAdapter(authService, userService, router)
	routers.NewPollHttpAdapter(pollRepo, voteRepo, cld, router)
	routers.NewUserHttpAdapter(userService, router)
	routers.NewStaticHttpAdapter(router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	logger.Error.Fatal(server.ListenAndServe())
}

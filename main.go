package main

import (
	"log"
	"net/http"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/http/routers"
	"github.com/bythecover/backend/persistence"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/goloop/env"
)

func main() {
	if err := env.Update(".env"); err != nil {
		log.Fatalln(err)
	}

	cld, err := cloudinary.New()
	if err != nil {
		log.Fatalln("Unable to instantiate cloudinary: ", err.Error())
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
	routers.NewCallbackHttpAdapter(authService, userService, router)
	routers.NewPollHttpAdapter(pollService, cld, router)
	routers.NewUserHttpAdapter(userService, router)
	routers.NewStaticHttpAdapter(router)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

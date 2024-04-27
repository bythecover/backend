package main

import (
	http_adapter "bythecover/backend/internal/adapters/http"
	"bythecover/backend/internal/adapters/persistence"
	"bythecover/backend/internal/core/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dbConnection := persistence.NewPostgresDatabase()

	userRepo := persistence.NewUserPostgresRepository(dbConnection)
	userService := services.NewUserService(userRepo)
	userHandler := http_adapter.NewUserHttpHandler(userService)

	voteRepo := persistence.NewVotePostgresRepository(dbConnection)

	pollRepo := persistence.NewPollPostgresRepository(dbConnection)
	pollService := services.NewPollService(pollRepo, voteRepo)
	pollAdapter := http_adapter.NewPollHttpHandler(pollService)

	testHandler := http_adapter.TestHttpHandler{}

	router := http.NewServeMux()

	userHandler.RegisterRoutes(router)
	pollAdapter.RegisterRoutes(router)
	testHandler.RegisterRoutes(router)

	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/assets"))))
	middlewareStack := http_adapter.CreateStack(http_adapter.AllowCors, http_adapter.Logger)

	server := http.Server{
		Handler: middlewareStack(router),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

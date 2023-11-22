package main

import (
	"bythecover/backend/internal/adapters/handler"
	"bythecover/backend/internal/adapters/repository"
	"bythecover/backend/internal/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	repo, _ := repository.NewUserPostgresRepository()

	userService := services.NewUserService(repo)
	userHandler := handler.NewUserHttpHandler(userService)

	route := gin.Default()
	userHandler.RegisterRoutes(route)
	route.Run(":8080")
}
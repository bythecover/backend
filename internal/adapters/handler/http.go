package handler

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"log"

	"github.com/gin-gonic/gin"
)

type UserHttpHandler struct {
	service ports.UserService
}

func NewUserHttpHandler(service ports.UserService) UserHttpHandler {
	return UserHttpHandler{
		service,
	}
}

func (adapter UserHttpHandler) RegisterRoutes(route *gin.Engine) {
	route.POST("/createUser", func (c *gin.Context) {
		var person domain.User
		if err := c.Bind(&person); err != nil {
			log.Print(err)
			return
		}

		adapter.service.Create(person)
	})

	route.GET("/users", func (c *gin.Context) {
		people, err := adapter.service.GetAll()

		if err != nil {
			c.AbortWithStatus(404)
		}

		for _, v := range people {
			log.Print(v)
		}

		c.JSON(200, people)
	})
}
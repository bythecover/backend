package handler

import (
	"bythecover/backend/internal/core/ports"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHttpHandler struct {
	service ports.UserService
}

func NewUserHttpHandler(service ports.UserService) userHttpHandler {
	return userHttpHandler{
		service,
	}
}

func (adapter userHttpHandler) RegisterRoutes(route *gin.Engine) {
	route.POST("/createUser", func (c *gin.Context) {
		var person ports.UserReq
		if err := c.Bind(&person); err != nil {
			log.Print(err)
			return
		}

		err := adapter.service.Create(person)

		if err != nil {
			c.AbortWithStatusJSON(400, err.Error())
		}

	})

	route.GET("/users/:id", func (c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Print(err)
			c.AbortWithStatus(500)
		}

		user, err := adapter.service.GetUser(id, c)

		if err != nil {
			if err == ports.ErrUserNotFound {
				c.AbortWithStatus(404)
				return;
			} else {
				c.AbortWithStatus(500)
				return;
			}
		}

		c.JSON(200, user)
		return;
	})

	route.GET("/users/")
}
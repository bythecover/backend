package user_handler

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
		var person ports.UserResp
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
			} else {
				c.AbortWithStatus(500)
			}
		} else {
			c.JSON(200, user)
		}

	})

	route.GET("/users/")
}
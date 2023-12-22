package poll_handler

import (
	"bythecover/backend/internal/core/ports"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pollHttpHandler struct {
	service ports.PollService
}

func NewPollHttpHandler(svc ports.PollService) pollHttpHandler {
	return pollHttpHandler{
		svc,
	}
}

func (handler pollHttpHandler) RegisterRoutes(route *gin.Engine) {
	route.GET("/api/polls/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Print(err)
			c.AbortWithStatus(500)
		} else {
			poll, err := handler.service.GetById(c, id)

			// TODO: handle errors better here
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(400)
			} else {
				c.JSON(200, poll)
			}
		}
	})
}

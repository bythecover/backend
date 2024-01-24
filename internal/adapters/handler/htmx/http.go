package htmx_handler

import (
	"bythecover/backend/internal/core/ports"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type htmxHttpHandler struct {
	htmxSvc ports.HtmxService
	pollSvc ports.PollService
}

func NewHtmxHttpHandler(htmxSvc ports.HtmxService, pollSvc ports.PollService) htmxHttpHandler {
	return htmxHttpHandler{
		htmxSvc,
		pollSvc,
	}
}

func (handler htmxHttpHandler) RegisterRoutes(route *gin.Engine) {
	route.GET("/vote/:id", func(c *gin.Context) {
		// convert string to number
		id, _ := strconv.Atoi(c.Param("id"))

		poll, _ := handler.pollSvc.GetById(c, id)

		handler.htmxSvc.VotePage(poll, c)
	})

	route.POST("/vote/submit", func(c *gin.Context) {
		// accept submission and return success/fail dialog
		content := c.PostForm("Option 2")
		log.Print(content)
		handler.htmxSvc.SubmitVote(c)
	})
}

package htmx_handler

import (
	"bythecover/backend/internal/adapters/handler/htmx/templates"
	"bythecover/backend/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type htmxHandler struct {
	pollService ports.PollService
}

func NewHtmxHandler(svc ports.PollService) htmxHandler {
	return htmxHandler{
		pollService: svc,
	}
}

func (handler htmxHandler) RegisterRoutes(route *gin.Engine) {

	route.POST("/clicked/:id", func(c *gin.Context) {
		templates.Clicked().Render(c, c.Writer)

	})

	route.GET("/newPage", func(c *gin.Context) {
		templates.NewPage().Render(c, c.Writer)
	})

	route.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

}

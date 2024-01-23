package htmx

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/services/htmx/components"
	"bythecover/backend/internal/core/services/htmx/pages"

	"github.com/gin-gonic/gin"
)

type htmxService struct {
}

func NewHtmxService() htmxService {
	return htmxService{}
}

func (service htmxService) VotePage(poll domain.Poll, c *gin.Context) error {
	c.Header("Content-Type", "text/html")
	pages.VotePage(poll).Render(c, c.Writer)
	return nil
}

func (service htmxService) SubmitVote(c *gin.Context) error {
	components.Dialog().Render(c, c.Writer)
	return nil
}

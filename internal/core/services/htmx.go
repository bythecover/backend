package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"bythecover/backend/internal/core/services/templates/components"
	"bythecover/backend/internal/core/services/templates/pages"

	"github.com/gin-gonic/gin"
)

type htmxService struct {
	voteService ports.VoteService
}

func NewHtmxService(voteService ports.VoteService) htmxService {
	return htmxService{
		voteService,
	}
}

func (service htmxService) VotePage(poll domain.Poll, c *gin.Context) error {
	c.Header("Content-Type", "text/html")
	pages.VotePage(poll).Render(c, c.Writer)
	return nil
}

func (service htmxService) SubmitVote(c *gin.Context) error {
	err := service.voteService.SubmitVote(c, 1)

	if err != nil {
		components.Dialog(err).Render(c, c.Writer)
		return nil
	}

	components.Dialog(nil).Render(c, c.Writer)
	return nil
}

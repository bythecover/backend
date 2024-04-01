package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"bythecover/backend/internal/core/services/templates/components"
	"bythecover/backend/internal/core/services/templates/pages"

	"github.com/a-h/templ"
)

type htmxService struct {
	voteService ports.VoteService
}

func NewHtmxService(voteService ports.VoteService) htmxService {
	return htmxService{
		voteService,
	}
}

func (service htmxService) VotePage(poll domain.Poll) templ.Component {
	return pages.VotePage(poll)
}

func (service htmxService) SubmitVote() templ.Component {
	err := service.voteService.SubmitVote(1)

	if err != nil {
		return components.Dialog(err)
	}

	return components.Dialog(nil)
}

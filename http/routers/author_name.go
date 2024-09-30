package routers

import (
	"net/http"
	"strconv"

	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/pages"
)

type authorHttpAdapter struct {
	pollRepo model.PollRepo
}

func NewAuthorHttpAdapter(router *http.ServeMux, pollRepo model.PollRepo) authorHttpAdapter {
	adapter := authorHttpAdapter{pollRepo}
	adapter.registerRoutes(router)
	return adapter
}

func (adapter *authorHttpAdapter) registerRoutes(router *http.ServeMux) {
	isAuthorized := middleware.IsAuthorizedAsAuthor()
	router.Handle("GET /a/{authorName}", isAuthorized(http.HandlerFunc(adapter.getAuthorPage)))
	router.Handle("PUT /a/{authorName}/{bookId}", isAuthorized(http.HandlerFunc(adapter.finalizePoll)))
}

func (adapter *authorHttpAdapter) getAuthorPage(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		logger.Warn.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authorId := r.PathValue("authorName")

	if authorId != session.Profile.UserId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	polls, err := adapter.pollRepo.GetCreatedBy(authorId)

	if err != nil {
		logger.Error.Println(err)
	}

	pages.AuthorPage(session, polls).Render(r.Context(), w)
}

func (adapter *authorHttpAdapter) finalizePoll(w http.ResponseWriter, r *http.Request) {
	pollId, err := strconv.Atoi(r.PathValue("bookId"))

	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Check to see if author has permissions to expire this poll
	adapter.pollRepo.ExpirePoll(pollId)

	w.WriteHeader(http.StatusOK)
}

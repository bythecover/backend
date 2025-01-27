package routers

import (
	"net/http"
	"strconv"

	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/pages"
)

type authorHttpAdapter struct {
	pollRepo model.PollRepo
}

func RegisterAuthorRoutes(router *http.ServeMux, pollRepo model.PollRepo) {
	adapter := authorHttpAdapter{pollRepo}
	isAuthorized := IsAuthorizedAsAuthor()
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
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		logger.Error.Println(err)
	}

	pollId, err := strconv.Atoi(r.PathValue("bookId"))

	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = adapter.pollRepo.ExpirePoll(pollId, session.Profile.UserId)
	if err != nil {
		logger.Warn.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Refresh", "true")
	w.WriteHeader(http.StatusOK)
}

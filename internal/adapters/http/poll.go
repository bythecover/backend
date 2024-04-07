package http

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"bythecover/backend/internal/core/services/templates/components"
	"bythecover/backend/internal/core/services/templates/pages"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

type pollHttpHandler struct {
	poll ports.PollService
}

func NewPollHttpHandler(poll ports.PollService) pollHttpHandler {
	return pollHttpHandler{
		poll,
	}
}

func (adapter pollHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		// convert string to number
		id, _ := strconv.Atoi(r.PathValue("id"))
		poll, _ := adapter.poll.GetById(id)
		templ.Handler(pages.VotePage(poll)).ServeHTTP(w, r)
	})

	// handle vote submission
	router.HandleFunc("POST /polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		selectedId, _ := strconv.Atoi(r.PostFormValue("selection"))
		pollId, _ := strconv.Atoi(r.PathValue("id"))

		submission := domain.Vote{
			Selection:   selectedId,
			PollEventId: pollId,
		}

		err := adapter.poll.SubmitVote(submission)
		htmx := components.Dialog(nil)

		if err != nil {
			htmx = components.Dialog(err)
		}

		templ.Handler(htmx).ServeHTTP(w, r)
	})
}

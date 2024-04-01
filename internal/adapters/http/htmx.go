package http

import (
	"bythecover/backend/internal/core/ports"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
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

func (handler htmxHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /votes/{id}", func(w http.ResponseWriter, r *http.Request) {
		// convert string to number
		id, _ := strconv.Atoi(r.PathValue("id"))
		poll, _ := handler.pollSvc.GetById(id)
		templ.Handler(handler.htmxSvc.VotePage(poll)).ServeHTTP(w, r)
	})

	// handle vote submission
	router.HandleFunc("POST /votes/:id", func(w http.ResponseWriter, r *http.Request) {
		// accept submission and return success/fail dialog
		r.ParseForm()
		content := r.PostFormValue("selection")
		log.Print(content)
		templ.Handler(handler.htmxSvc.SubmitVote()).ServeHTTP(w, r)
	})
}

package http

import (
	"bythecover/backend/internal/core/ports"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type pollHttpHandler struct {
	service ports.PollService
}

func NewPollHttpHandler(svc ports.PollService) pollHttpHandler {
	return pollHttpHandler{
		svc,
	}
}

func (handler pollHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
		} else {
			poll, err := handler.service.GetById(id)

			// TODO: handle errors better here
			if err != nil {
				log.Print(err)
				w.WriteHeader(400)
			} else {
				data, _ := json.Marshal(poll)
				w.Write(data)
			}
		}
	})
}

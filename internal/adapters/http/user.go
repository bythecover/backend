package http

import (
	"bythecover/backend/internal/core/ports"
	"log"
	"strconv"

	"encoding/json"
	"io"
	"net/http"
)

type userHttpHandler struct {
	userService ports.UserService
}

func NewUserHttpHandler(userService ports.UserService) userHttpHandler {
	return userHttpHandler{
		userService,
	}
}

func decode[V any](r io.Reader, p V) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (adapter userHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		var person ports.UserResp
		err := decode(r.Body, &person)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = adapter.userService.Create(person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		user, err := adapter.userService.GetUser(id)

		if err != nil {
			if err == ports.ErrUserNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			user, _ := json.Marshal(user)
			w.Write(user)
		}

	})
}

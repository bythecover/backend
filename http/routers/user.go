package routers

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/services"
	"log"
	"strconv"

	"encoding/json"
	"io"
	"net/http"
)

type userHttpAdapter struct {
	userService services.UserService
}

func NewUserHttpAdapter(userService services.UserService, router *http.ServeMux) userHttpAdapter {
	adapter := userHttpAdapter{
		userService,
	}
	adapter.RegisterRoutes(router)
	return adapter
}

func decode[V any](r io.Reader, p V) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (adapter userHttpAdapter) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /user", adapter.createUser)
	router.HandleFunc("GET /user/{id}", adapter.getUser)
}

func (adapter userHttpAdapter) createUser(w http.ResponseWriter, r *http.Request) {
	var person model.UserResp
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
}

func (adapter userHttpAdapter) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	user, err := adapter.userService.GetUser(id)

	if err != nil {
		if err == model.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		user, _ := json.Marshal(user)
		w.Write(user)
	}

}

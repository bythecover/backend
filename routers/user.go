package routers

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"

	"encoding/json"
	"io"
	"net/http"
)

type userHttpAdapter struct {
	userRepo persistence.UserRepo
}

func NewUserHttpAdapter(userRepo persistence.UserRepo, router *http.ServeMux) userHttpAdapter {
	adapter := userHttpAdapter{
		userRepo,
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
}

func (adapter userHttpAdapter) createUser(w http.ResponseWriter, r *http.Request) {
	var person model.UserResp
	err := decode(r.Body, &person)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := model.NewUser(person.Id, person.FirstName, person.LastName, person.Email, person.Role)

	err = adapter.userRepo.Save(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

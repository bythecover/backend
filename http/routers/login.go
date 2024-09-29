package routers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/sessions"
)

type loginHttpAdapter struct {
	authenticator *authenticator.Authenticator
}

func NewLoginHttpAdapter(authenticator *authenticator.Authenticator, router *http.ServeMux) loginHttpAdapter {
	adapter := loginHttpAdapter{
		authenticator,
	}
	adapter.registerRoutes(router)
	return adapter
}

func (adapter loginHttpAdapter) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /login", adapter.loginHandler)
}

func (adapter loginHttpAdapter) loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err)
		return
	}

	if session.State != "" {
		logger.Warn.Println("session already has state")
		w.WriteHeader(http.StatusOK)
		return
	}

	state, _ := generateRandomState()
	session.State = state
	session.Save()
	http.Redirect(w, r, adapter.authenticator.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

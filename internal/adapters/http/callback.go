package http

import (
	"bythecover/backend/internal/core/services/authenticator"
	"bythecover/backend/internal/core/services/sessions"
	"log"
	"net/http"
)

type callbackHttpAdapter struct {
	authenticator *authenticator.Authenticator
}

func NewCallbackHttpAdapter(authenticator *authenticator.Authenticator) callbackHttpAdapter {
	return callbackHttpAdapter{
		authenticator,
	}
}

// The callback is the endpoint that the user is redirected to after authenticating with Auth0.
// Search Auth0 logs for "callback" to see more
func (adapter callbackHttpAdapter) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /callback", adapter.handler)
}

func (adapter callbackHttpAdapter) handler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.WithSession(r.Context())
	queryState := r.URL.Query().Get("state")
	requestHasValidState := queryState == session.State
	if !requestHasValidState {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryCode := r.URL.Query().Get("code")
	token, err := adapter.authenticator.Exchange(r.Context(), queryCode)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idToken, err := adapter.authenticator.VerifyIDToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Profile = profile
	session.AccessToken = token.AccessToken
	session.Save()

	log.Println(session)

	// Redirect to logged in page.
	http.Redirect(w, r, "/", http.StatusOK)
}

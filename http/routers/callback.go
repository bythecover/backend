package routers

import (
	"net/http"
	"time"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
)

type callbackHttpAdapter struct {
	authenticator *authenticator.Authenticator
	userService   services.UserService
}

func NewCallbackHttpAdapter(authenticator *authenticator.Authenticator, userService services.UserService, router *http.ServeMux) callbackHttpAdapter {
	adapter := callbackHttpAdapter{
		authenticator,
		userService,
	}

	adapter.RegisterRoutes(router)
	return adapter
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

	var profile sessions.Auth0User
	if err := idToken.Claims(&profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Profile = profile
	session.AccessToken = token.AccessToken
	session.Save()

	existingUser, err := adapter.userService.GetUser(session.Profile.UserId)
	if err != nil {
		currentTime := time.Now()
		adapter.userService.Create(model.UserResp{
			Id:        session.Profile.UserId,
			FirstName: session.Profile.Nickname,
			LastName:  session.Profile.Name,
			Email:     "dummy@email.com",
			Role:      "user",
			CreatedAt: &currentTime,
		})

		session.Profile.Role = "user"
	} else {
		session.Profile.Role = existingUser.Role
	}

	// Redirect to logged in page.
	http.Redirect(w, r, "/", http.StatusOK)
}

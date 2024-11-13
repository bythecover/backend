package routers

import (
	"net/http"

	"github.com/bythecover/backend/authenticator"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/pages"
)

type callbackHttpAdapter struct {
	authenticator *authenticator.Authenticator
	userRepo      persistence.UserRepo
}

func NewCallbackHttpAdapter(authenticator *authenticator.Authenticator, userRepo persistence.UserRepo, router *http.ServeMux) callbackHttpAdapter {
	adapter := callbackHttpAdapter{
		authenticator,
		userRepo,
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
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.FiveHundred().Render(r.Context(), w)
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

	existingUser, err := adapter.userRepo.GetUser(session.Profile.UserId)
	if err != nil {
		session.Profile.Role = "user"
		if err == model.ErrUserNotFound {
			user, _ := model.NewUser(session.Profile.UserId, session.Profile.Nickname, session.Profile.Name, "test@test.com", session.Profile.Role)

			err = adapter.userRepo.Save(user)
		}
	} else {
		session.Profile.Role = existingUser.Role
	}

	// Redirect to logged in page.
	http.Redirect(w, r, "/", http.StatusOK)
}

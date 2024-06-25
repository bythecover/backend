package sessions

import (
	"context"

	"github.com/google/uuid"
)

const SESSION_COOKIE_NAME = "sessionid"

// A Session stores information about a user's Session after they have logged in
type Session struct {
	store       SessionStore
	State       string
	Profile     Auth0User
	AccessToken string
	Id          uuid.UUID
}

type Auth0User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Picture  string `json:"picture"`
	UserId   string `json:"sub"`
	Role     string
}

// Creates a new Session
func New() *Session {
	return &Session{
		store: globalStore,
	}
}

// Persists the session to the store
func (session *Session) Save() uuid.UUID {
	return globalStore.Save(session)
}

// NewContext returns a context with the session added as a value
func NewContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, "session", session)
}

// FromContext will attempt to get a session from a context value
func FromContext(ctx context.Context) (*Session, error) {
	session, ok := ctx.Value("session").(*Session)

	if ok != true {
		return nil, ErrSessionTypeCastFailed
	}

	if session == nil {
		return nil, ErrSessionNotFound
	}

	return session, nil
}

// WithSession will check to see if a session exists in a context.
// If one does not exist, it will create one. It is the responsibility
// to the caller to add the session to the context (with sessions.NewContext) and save it to the store
// when appropriate.
func WithSession(ctx context.Context) (*Session, error) {
	var session *Session
	var err error

	if session, err = FromContext(ctx); err != nil {
		session = &Session{}
	}

	return session, err
}

// New implementations can be made of the Session Store interface as needed
// This sessions package provides some implementations but not all
type SessionStore interface {
	Save(*Session) uuid.UUID
	Get(string) (*Session, error)
}

var globalStore SessionStore

// Create store should be called on setup to allow for
func CreateStore(store SessionStore) {
	globalStore = store
}

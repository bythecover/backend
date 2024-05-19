package http

import "context"

// A Session stores information about a user's Session after they have logged in
type Session struct {
	State string
}

type SessionStore map[string]*Session

func FromContext(ctx context.Context) (*Session, bool) {
	session, ok := ctx.Value("session").(*Session)
	return session, ok
}

const SESSION_COOKIE_NAME = "sessionid"

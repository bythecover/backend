package sessions

import "errors"

var (
	ErrSessionNotFound       = errors.New("Session not found in Store")
	ErrSessionTypeCastFailed = errors.New("Failed attempt to cast Context Value to Session pointer")
)

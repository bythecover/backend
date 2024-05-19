// Package
package http

import (
	"context"
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
)

// A WrappedWriter exposes the status code to be able to print in the Logger
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &wrappedWriter{
			w,
			http.StatusOK,
		}
		next.ServeHTTP(writer, r)
		log.Println(writer.statusCode, r.Method, r.URL.Path)
	})
}

// HandlerWithSessionStore Creates a Middleware function that adds the session store
// to the context of the current request
func HandlerWithSession(store SessionStore) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var session *Session

			cookieIndex := slices.IndexFunc(r.Cookies(), func(c *http.Cookie) bool { return c.Name == SESSION_COOKIE_NAME })
			requestHasSessionCookie := cookieIndex >= 0

			if requestHasSessionCookie {
				// TODO: Handle when the store lookup fails
				log.Println("request has Session cookie")
				session = store[r.Cookies()[cookieIndex].Value]
			} else {
				session = &Session{
					State: "",
				}

				// save new session to store
				sessionId := uuid.New().String()
				store[sessionId] = session

				// send sessionid cookie to client
				sessionCookie := http.Cookie{
					Name:     "sessionid",
					Value:    sessionId,
					Secure:   true,
					HttpOnly: true,
					MaxAge:   0,
				}
				http.SetCookie(w, &sessionCookie)
			}

			// add the session object to the context
			log.Println("setting session to the following")
			log.Println(session)
			newContext := context.WithValue(r.Context(), "session", session)
			newRequest := r.WithContext(newContext)
			next.ServeHTTP(w, newRequest)
		})
	}
}

type Middleware func(http.Handler) http.Handler

// Create stack takes a list of middleware to run in order
func CreateStack(handlers ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			next = handlers[i](next)
		}

		return next
	}
}

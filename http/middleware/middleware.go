package middleware

import (
	"github.com/bythecover/backend/sessions"
	"log"
	"net/http"
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
func HandlerWithSession(store sessions.SessionStore) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var session *sessions.Session

			cookie, err := r.Cookie(sessions.SESSION_COOKIE_NAME)

			if err == nil {
				sessionId := cookie.Value
				session, err = store.Get(sessionId)

				if err != nil {
					session = sessions.New()
					addNewSessionToCookie(w, session)
				}
			} else {
				session = sessions.New()
				addNewSessionToCookie(w, session)
			}

			// Add the session to the context
			newContext := sessions.NewContext(r.Context(), session)
			next.ServeHTTP(w, r.WithContext(newContext))
		})
	}
}

// Create a new Session and save its ID to the cookie and send the cookie as a response
func addNewSessionToCookie(w http.ResponseWriter, session *sessions.Session) {
	sessionId := session.Save().String()
	sessionCookie := http.Cookie{
		Name:     "sessionid",
		Value:    sessionId,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   0,
	}

	http.SetCookie(w, &sessionCookie)
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

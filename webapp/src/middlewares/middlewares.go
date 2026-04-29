package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger writes request info in the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verify cookies existence
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next(w, r)
	}
}

package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout removes the authentication cookie and redirects the user to the login page
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

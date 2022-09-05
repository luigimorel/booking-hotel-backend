package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	_, ok := session.Values["user"]

	if !ok {
		http.Redirect(w, r, os.Getenv("BASE_URL"), http.StatusTemporaryRedirect)

	}
}

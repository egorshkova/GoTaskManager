package app

import (
	u "ServerApp/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)//http 404 code
		u.Respond(w, u.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
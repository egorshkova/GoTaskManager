package app

import (
	u "ServerApp/utils"
	"ServerApp/models"
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //list of endpoints for which authentification is not required
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		//check if authentification is needed for requestPath
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenFromHeader := r.Header.Get("Authorization")

		if tokenFromHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)	//http code 403
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		//check correctness of token, required format: `Bearer {token-body}`
		splittedTokenFromHeader := strings.Split(tokenFromHeader, " ")
		if len(splittedTokenFromHeader) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden) //http code 403
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splittedTokenFromHeader[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)//http code 403
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden) //http code 403
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		fmt.Sprintf("User %d", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
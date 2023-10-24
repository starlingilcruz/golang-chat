package middlewares

import (
	"strings"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/starlingilcruz/golang-chat/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !userIsAuthenticated(w, r) {
				http.Error(w, "Authorization header is missing", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
	})
}

func userIsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	authz := r.Header.Get("authorization")

	if authz == "" {
		return false
	}

	parts := strings.Split(authz, " ")

	if len(parts) != 2 {
		http.Error(w, "Malformed token", http.StatusForbidden)
		return false
	}

	tokenPart := parts[1]

	token, err := utils.ParseJWT(tokenPart)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return false
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		http.Error(w, "Token is not valid", http.StatusForbidden)
		return false
	}

	return true
}

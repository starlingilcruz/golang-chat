package middlewares

import (
	"net/http"
)


func Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO implement 

		next(w, r)
	}
}
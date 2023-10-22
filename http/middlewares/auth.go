package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)


func Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		fmt.Println(authHeader)


		// read basic auth information
		usr, _, ok := r.BasicAuth()
		fmt.Printf("User %s logged in.", usr)
		// if there is no basic auth (no matter which credentials)
		if !ok {
			errMsg := "Authentication error!"
			// return a 403 forbidden
			http.Error(w, errMsg, http.StatusForbidden)
			fmt.Println(errMsg)

			// stop processing route
			return
		}

		// let's assume we check the user against a database to get
		// his admin-right and put this to the request context
		// context.Set(r, "isAdmin", true)

		// else continue processing
		fmt.Printf("User %s logged in.", usr)
		next(w, r)
	}
}
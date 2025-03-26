package middleware

import "net/http"

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorized", "true")
		next(w, r)
	}
}

package middleware

import "net/http"

func SetHeaderResponse(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", "My custom header")
		next(w, r)
	}
}

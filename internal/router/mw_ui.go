package router

import (
	"net/http"
)

func MwUi(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("UI", "true")
		next.ServeHTTP(w, r)
	})
}

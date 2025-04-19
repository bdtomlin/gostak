package router

import (
	"context"
	"net/http"
)

func MwEnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path", r.URL.Path)
		w.Header().Add("EnsureAdmin", "TRUE")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

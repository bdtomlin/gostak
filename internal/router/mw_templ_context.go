package router

import (
	"context"
	"net/http"
)

func MwTemplContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path", r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

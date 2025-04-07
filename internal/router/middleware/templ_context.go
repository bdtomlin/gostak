package middleware

import (
	"context"
	"net/http"
)

func TemplContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path", r.URL.Path)
		next(w, r.WithContext(ctx))
	}
}

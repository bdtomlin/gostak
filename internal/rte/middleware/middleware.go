package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(middlewares ...Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

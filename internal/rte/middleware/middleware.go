package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Chain struct {
	MwChain func(http.HandlerFunc) http.HandlerFunc
}

func NewChain(middlewares ...Middleware) Chain {
	return Chain{
		MwChain: makeChain(middlewares),
	}
}

func (c Chain) HandleFunc(path string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(path, c.MwChain(handlerFunc))
}

func makeChain(middlewares []Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

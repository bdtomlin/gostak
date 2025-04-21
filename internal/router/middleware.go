package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MwChain struct {
	Mux        *http.ServeMux
	Middleware Middleware
}

func (mwc MwChain) Handle(pattern string, handler http.Handler) {
	mwc.Mux.Handle(pattern, mwc.Middleware(handler))
}

func (mwc MwChain) HandleFunc(pattern string, hf http.HandlerFunc) {
	mwc.Handle(pattern, http.HandlerFunc(hf))
}

func newMiddlewareChain(mux *http.ServeMux, mws ...Middleware) MwChain {
	mw := func(next http.Handler) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			next = mws[i](next)
		}
		return next
	}
	return MwChain{
		Mux:        mux,
		Middleware: mw,
	}
}

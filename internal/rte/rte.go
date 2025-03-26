package rte

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handlers"
	"github.com/bdtomlin/gostak/internal/rte/middleware"
)

func handleFunc(path string, handler http.HandlerFunc, chain func(http.HandlerFunc) http.HandlerFunc) {
	http.HandleFunc(path, chain(handler))
}

func Route() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	browser := middleware.NewChain(middleware.Html, middleware.SetHeaderResponse, middleware.Gzip)
	browser.HandleFunc("/", handlers.GetIndex)
	browser.HandleFunc("GET /users/{id}", handlers.GetUser)

	browserAuth := middleware.NewChain(browser.MwChain, middleware.Auth)
	browserAuth.HandleFunc("GET /authors", handlers.GetAuthors)
}

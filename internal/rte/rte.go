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

	browser := middleware.Chain(middleware.Html, middleware.SetHeaderResponse, middleware.Gzip)
	handleFunc("/", handlers.GetIndex, browser)
	handleFunc("GET /users/{id}", handlers.GetUser, browser)

	browserAuth := middleware.Chain(browser, middleware.Auth)
	handleFunc("GET /authors", handlers.GetAuthors, browserAuth)
}

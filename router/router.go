package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/handlers"
	"github.com/bdtomlin/gostak/router/middleware"
)

func handleFunc(path string, handler http.HandlerFunc, chain func(http.HandlerFunc) http.HandlerFunc) {
	http.HandleFunc(path, chain(handler))
}

func Route() {
	browser := middleware.Chain(middleware.Html, middleware.SetHeaderResponse, middleware.Gzip)
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	handleFunc("/", handlers.GetIndex, browser)
	handleFunc("GET /users/{id}", handlers.GetUser, browser)
	browserAuth := middleware.Chain(browser, middleware.Auth)
	handleFunc("GET /authors", handlers.GetAuthors, browserAuth)
}

package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
	"github.com/bdtomlin/gostak/internal/router/middleware"
)

func handleFunc(path string, handler http.HandlerFunc, chain func(http.HandlerFunc) http.HandlerFunc) {
	http.HandleFunc(path, chain(handler))
}

func Route() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	browser := middleware.NewChain(
		middleware.Html,
		middleware.SetHeaderResponse,
		middleware.Gzip,
		middleware.TemplContext,
	)
	browser.HandleFunc("/", handler.GetIndex)
	browser.HandleFunc("GET /users/{id}", handler.GetUser)

	// UI
	browser.HandleFunc("/ui/sign-up", handler.UiSignUp)
	browser.HandleFunc("/ui/sign-in", handler.UiSignIn)
	browser.HandleFunc("/ui/mailers", handler.UiListMailers)

	browserAuth := middleware.NewChain(browser.MwChain, middleware.Auth)
	browserAuth.HandleFunc("GET /authors", handler.GetAuthors)
}

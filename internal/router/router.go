package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
	"github.com/bdtomlin/gostak/internal/router/middleware"
)

func NewRouter() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/assets"))
	mux.Handle("/web/assets/", http.StripPrefix("/web/assets", fs))
}

func Route() {
	fs := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/web/assets/", http.StripPrefix("/web/assets", fs))

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
	browserAuth.HandleFunc("GET /users", handler.ListUsers)
}

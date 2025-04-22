package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/assets/"))
	mux.Handle("/web/assets/", http.StripPrefix("/web/assets/", fs))

	bChain := newMiddlewareChain(mux, MwLogger, MwTemplContext, MwGzip, MwCsrf)
	bChain.HandleFunc("/", handler.Index)
	bChain.HandleFunc("GET /users/{id}", handler.GetUser)
	bChain.HandleFunc("GET /users", handler.ListUsers)
	bChain.HandleFunc("GET /set-cookie", handler.SetCookie)
	bChain.HandleFunc("GET /show-cookie", handler.ShowCookie)

	// UI
	uiChain := newMiddlewareChain(mux, bChain.Middleware, MwUi)
	uiChain.HandleFunc("GET /ui/sign-up", handler.UiSignUp)
	uiChain.HandleFunc("GET /ui/sign-in", handler.UiSignIn)
	uiChain.HandleFunc("POST /ui/sign-in", handler.UiSignIn)
	uiChain.HandleFunc("PUT /ui/sign-in", handler.UiSignIn)
	uiChain.HandleFunc("GET /ui/mailers", handler.UiListMailers)

	return mux
}

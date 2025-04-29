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
	bChain.HandleFunc("GET /sign-up", handler.GetSignUp)
	bChain.HandleFunc("POST /sign-up", handler.PostSignUp)
	bChain.HandleFunc("GET /sign-in", handler.GetSignIn)
	bChain.HandleFunc("POST /sign-in", handler.PostSignIn)

	// UI
	uiChain := newMiddlewareChain(mux, bChain.Middleware, MwUi)
	uiChain.HandleFunc("GET /ui/mailers", handler.UiListMailers)

	return mux
}

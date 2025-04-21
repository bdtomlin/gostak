package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/assets/"))
	mux.Handle("/web/assets/", http.StripPrefix("/web/assets/", fs))

	bmw := chainMiddleware(MwTemplContext, MwGzip, MwLogger)
	mux.Handle("/", bmw(handler.Index{}))
	mux.Handle("GET /users/{id}", bmw(handler.GetUser{}))
	mux.Handle("GET /users", bmw(handler.ListUsers{}))

	// UI
	uimw := chainMiddleware(MwTemplContext, MwGzip, MwLogger, MwUi)
	mux.Handle("GET /ui/sign-up", uimw(handler.UiSignUp{}))
	mux.Handle("GET /ui/sign-in", uimw(handler.UiSignIn{}))
	mux.Handle("GET /ui/mailers", uimw(handler.UiListMailers{}))

	return mux
}

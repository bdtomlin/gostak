package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/assets"))
	mux.Handle("/web/assets/", http.StripPrefix("/web/assets", fs))
	mux.Handle("/", handler.NewIndexHandler())
	mux.Handle("GET /users/{id}", handler.NewGetUser())
	// UI
	mux.HandleFunc("/ui/sign-up", handler.UiSignUp)
	mux.HandleFunc("/ui/sign-in", handler.UiSignIn)
	mux.HandleFunc("/ui/mailers", handler.UiListMailers)

	return MwTemplContext(MwGzip(MwLogger(mux)))
}

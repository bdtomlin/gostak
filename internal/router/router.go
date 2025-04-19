package router

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/handler"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/assets"))
	router.Handle("/web/assets/", http.StripPrefix("/web/assets", fs))

	browserMw := NewMiddleware(MwTemplContext, MwGzip, MwLogger)
	pubrouter := http.NewServeMux()
	router.Handle("/", browserMw(pubrouter))
	pubrouter.Handle("/", handler.NewIndexHandler())
	pubrouter.Handle("GET /users/{id}", handler.NewGetUser())
	pubrouter.Handle("/users", handler.NewListUsers())

	// UI
	uiMw := NewMiddleware(browserMw, MwUi)
	uirouter := http.NewServeMux()
	router.Handle("/ui/", uiMw(uirouter))
	uirouter.HandleFunc("/ui/sign-up", handler.UiSignUp)
	uirouter.HandleFunc("/ui/sign-in", handler.UiSignIn)
	uirouter.HandleFunc("/ui/mailers", handler.UiListMailers)

	return router
}

package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

type Error404 struct{}

func (e Error404) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	page.Error404().Render(r.Context(), w)
}

func NewError404() Error404 {
	return Error404{}
}

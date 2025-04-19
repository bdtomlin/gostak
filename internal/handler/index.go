package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

type Index struct{}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Return 404 for anything other than "/"
		NewError404().ServeHTTP(w, r)
		return
	}
	page.Index().Render(r.Context(), w)
}

func NewIndexHandler() Index {
	return Index{}
}

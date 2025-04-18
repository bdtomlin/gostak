package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Return 404 for anything other than "/"
		Error(w, r, http.StatusNotFound, "")
		return
	}
	page.Index().Render(r.Context(), w)
}

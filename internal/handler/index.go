package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/htm/pages"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Return 404 for anything other than "/"
		Error(w, r, http.StatusNotFound)
		return
	}
	pages.Index().Render(r.Context(), w)
}

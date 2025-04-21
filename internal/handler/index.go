package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Return 404 for anything other than "/"
		Error404(w, r)
		return
	}
	page.Index().Render(r.Context(), w)
}

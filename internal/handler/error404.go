package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	page.Error404().Render(r.Context(), w)
}

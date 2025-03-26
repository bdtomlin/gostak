package handlers

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/htm/pages"
)

func Error(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		pages.Error(status).Render(r.Context(), w)
	}
}

package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/pages"
)

func Error(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		pages.Error(status).Render(r.Context(), w)
	}
}

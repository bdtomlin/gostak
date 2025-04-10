package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/pages"
)

func Error(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		pages.Error(status, message).Render(r.Context(), w)
	}
}

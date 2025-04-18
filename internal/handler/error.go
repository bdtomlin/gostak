package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func Error(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		page.Error(status, message).Render(r.Context(), w)
	}
}

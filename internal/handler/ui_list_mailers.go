package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/pages"
)

func UiListMailers(w http.ResponseWriter, r *http.Request) {
	pages.UiListMailers().Render(r.Context(), w)
}

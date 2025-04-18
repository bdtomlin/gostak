package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func UiListMailers(w http.ResponseWriter, r *http.Request) {
	page.UiListMailers().Render(r.Context(), w)
}

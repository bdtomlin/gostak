package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func UiSignIn(w http.ResponseWriter, r *http.Request) {
	page.UiSignIn().Render(r.Context(), w)
}

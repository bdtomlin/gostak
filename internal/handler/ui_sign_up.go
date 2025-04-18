package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func UiSignUp(w http.ResponseWriter, r *http.Request) {
	page.UiSignUp().Render(r.Context(), w)
}

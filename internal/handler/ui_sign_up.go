package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/pages"
)

func UiSignUp(w http.ResponseWriter, r *http.Request) {
	pages.UiSignUp().Render(r.Context(), w)
}

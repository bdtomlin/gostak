package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/htm/pages"
)

func UiSignIn(w http.ResponseWriter, r *http.Request) {
	pages.UiSignIn().Render(r.Context(), w)
}

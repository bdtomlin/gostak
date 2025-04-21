package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

type UiSignIn struct{}

func (u UiSignIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.UiSignIn().Render(r.Context(), w)

}

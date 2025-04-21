package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

type UiSignUp struct{}

func (u UiSignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.UiSignUp().Render(r.Context(), w)

}

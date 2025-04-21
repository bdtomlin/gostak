package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

type UiListMailers struct{}

func (u UiListMailers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.UiListMailers().Render(r.Context(), w)

}

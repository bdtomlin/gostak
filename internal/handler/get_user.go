package handler

import (
	"net/http"
	"strconv"

	"github.com/bdtomlin/gostak/web/page"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	page.User(id).Render(r.Context(), w)
}

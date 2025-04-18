package handler

import (
	"net/http"
	"strconv"

	"github.com/bdtomlin/gostak/web/page"
)

type getUserParams struct{}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	page.User(id).Render(r.Context(), w)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/bdtomlin/gostak/internal/htm/pages"
)

type getUserParams struct{}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	pages.User(id).Render(r.Context(), w)
}

package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/web/page"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.ListUsers()
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	page.ListUsers(users).Render(r.Context(), w)
}

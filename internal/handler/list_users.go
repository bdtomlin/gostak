package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/bdtomlin/gostak/web/page"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := model.UserRepo{}
	users, err := userRepo.ListUsers()
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	page.ListUsers(users).Render(r.Context(), w)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/web/page"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.ListUsers()
	if err != nil {
		// Return 500 for server errors
		errMsg := fmt.Sprintf("%v", err)
		Error500(w, r, errMsg)
		return
	}

	page.ListUsers(users).Render(r.Context(), w)
}

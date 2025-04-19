package handler

import (
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/web/page"
)

type ListUsers struct{}

func (lu ListUsers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := repo.ListUsers()
	if err != nil {
		// Return 500 for server errors
		errMsg := fmt.Sprintf("%v", err)
		NewError500(errMsg).ServeHTTP(w, r)
		return
	}

	page.ListUsers(users).Render(r.Context(), w)
}

func NewListUsers() ListUsers {
	return ListUsers{}
}

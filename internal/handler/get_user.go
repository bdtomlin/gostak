package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bdtomlin/gostak/web/page"
)

type GetUser struct{}

func (gu GetUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		NewError500(fmt.Sprintf("%v", err)).ServeHTTP(w, r)
		return
	}

	page.User(id).Render(r.Context(), w)
}

func NewGetUser() GetUser {
	return GetUser{}
}

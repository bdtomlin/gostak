package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	page.SignUp().Render(r.Context(), w)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
}

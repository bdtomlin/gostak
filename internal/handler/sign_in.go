package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/component"
	"github.com/bdtomlin/gostak/web/page"
)

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	page.SignIn().Render(r.Context(), w)
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	page.SignInForm("one", "two").Render(r.Context(), w)
	component.Notification("success", "some message", "some description").Render(r.Context(), w)
}

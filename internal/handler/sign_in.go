package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/component"
	"github.com/bdtomlin/gostak/web/form"
	"github.com/bdtomlin/gostak/web/page"
)

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignIn()
	page.SignIn(f).Render(r.Context(), w)
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignIn()

	err := decodeParams(f, r)
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	session, err := f.Submit()
	if err != nil {
		page.SignInForm(f).Render(r.Context(), w)
		component.Notifications(component.Notification{Variant: "error", Message: err.Error(), Description: ""}).Render(r.Context(), w)
		return
	}

	setSessionCookie(w, session)
	redirect(w, r, "/")
}

package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/form"
	"github.com/bdtomlin/gostak/web/page"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignUp()
	page.SignUp(f).Render(r.Context(), w)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignUp()

	err := decodeParams(f, r)
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	session, err := f.Submit()
	if err != nil {
		page.SignUpForm(f).Render(r.Context(), w)
		return
	}
	setSessionCookie(w, session)
	redirect(w, r, "/")
}

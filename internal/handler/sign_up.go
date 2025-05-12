package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/form"
	"github.com/bdtomlin/gostak/web/page"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	suf := form.NewSignUp()
	page.SignUp(*suf).Render(r.Context(), w)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	suf := form.NewSignUp()

	err := DecodeParams(suf, r)
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	if suf.IsValid() {
		// do sign up stuff
	} else {
		page.SignUpForm(*suf).Render(r.Context(), w)
	}
}

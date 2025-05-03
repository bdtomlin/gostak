package handler

import (
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/web/form"
	"github.com/bdtomlin/gostak/web/page"
	"github.com/gorilla/schema"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	suf := form.NewSignUp()
	page.SignUp(*suf).Render(r.Context(), w)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	var decoder = schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	suf := form.NewSignUp()

	err = decoder.Decode(suf, r.PostForm)
	if err != nil {
		fmt.Println(err)
	}

	suf.Validate()
	page.SignUp(*suf).Render(r.Context(), w)
}

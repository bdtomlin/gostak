package handler

import (
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/web/form"
	"github.com/bdtomlin/gostak/web/page"
)

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignIn()
	page.SignIn(f).Render(r.Context(), w)
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	f := form.NewSignIn()

	err := DecodeParams(f, r)
	if err != nil {
		RenderError500(err, w, r)
		return
	}

	user, err := f.Submit()
	if err != nil {
		page.SignInForm(f).Render(r.Context(), w)
		return
	}
	fmt.Println(user)
	Redirect(w, r, "/")

	// component.Notification("success", "some message", "some description").Render(r.Context(), w)
	// component.Notification("error", "some message", "some description").Render(r.Context(), w)
	// component.Notification("warning", "some message", "some description").Render(r.Context(), w)
	// component.Notification("info", "some message", "some description").Render(r.Context(), w)
	// component.Notifications(
	// 	component.Notification{Variant: "success", Message: "some message", Description: "some description"},
	// 	component.Notification{Variant: "error", Message: "some message", Description: "some description"},
	// 	component.Notification{Variant: "warning", Message: "some message", Description: "some description"},
	// 	component.Notification{Variant: "info", Message: "some message", Description: "some description"},
	// ).Render(r.Context(), w)
}

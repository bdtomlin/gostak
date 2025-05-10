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
	// component.Notification("success", "some message", "some description").Render(r.Context(), w)
	// component.Notification("error", "some message", "some description").Render(r.Context(), w)
	// component.Notification("warning", "some message", "some description").Render(r.Context(), w)
	// component.Notification("info", "some message", "some description").Render(r.Context(), w)
	component.Notifications(
		component.Notification{Variant: "success", Message: "some message", Description: "some description"},
		component.Notification{Variant: "error", Message: "some message", Description: "some description"},
		component.Notification{Variant: "warning", Message: "some message", Description: "some description"},
		component.Notification{Variant: "info", Message: "some message", Description: "some description"},
	).Render(r.Context(), w)
}

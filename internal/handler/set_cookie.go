package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "theCookie",
		Value:    "theValue",
		Path:     "/",  // path is optional, could be usefull for something like "/admin"
		HttpOnly: true, // don't allow js to access cookie for csrf security
	}
	http.SetCookie(w, &cookie)

	page.SetCookie().Render(r.Context(), w)
}

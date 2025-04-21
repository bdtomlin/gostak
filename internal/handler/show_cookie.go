package handler

import (
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/web/page"
)

func ShowCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("theCookie")
	if err != nil {
		fmt.Fprint(w, "The cookie couldn't be read.")
		return
	}

	page.ShowCookie(cookie.Value).Render(r.Context(), w)
}

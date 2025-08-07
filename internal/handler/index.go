package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/bdtomlin/gostak/web/page"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Return 404 for anything other than "/"
		Error404(w, r)
		return
	}

	uid, sid, err := readSessionCookie(r)
	if err != nil {
		slog.Error(err.Error())
	} else {
		sessionRepo := model.SessionRepo{}
		// user, err := sessionRepo.GetSessionUser(uid, sid)
		// if err != nil {
		// 	slog.Error(err.Error())
		// }
		// fmt.Printf("%+v", user)
		// sessionRepo.DeleteSession(uid, sid)
		fmt.Println(sid)
		sessionRepo.DeleteAllUserSessions(uid)
	}
	page.Index().Render(r.Context(), w)
}

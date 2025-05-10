package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/bdtomlin/gostak/web/page"
)

type Error500 struct {
	Message string
}

func (e Error500) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stderr, "%v\n", e.Message)
	w.WriteHeader(http.StatusInternalServerError)
	page.Error500(e.Message).Render(r.Context(), w)
}

func RenderError500(err error, w http.ResponseWriter, r *http.Request) {
	message := fmt.Errorf("500 error: %w", err).Error()
	slog.Error(message)
	Error500{Message: message}.ServeHTTP(w, r)
}

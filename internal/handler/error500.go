package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bdtomlin/gostak/web/page"
)

func Error500(w http.ResponseWriter, r *http.Request, message string) {
	// slog.Error(message)
	fmt.Fprintf(os.Stderr, "%v", message)
	w.WriteHeader(500)
	page.Error500(message).Render(r.Context(), w)
}

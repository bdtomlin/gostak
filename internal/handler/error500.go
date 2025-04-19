package handler

import (
	"fmt"
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

func NewError500(message string) Error500 {
	return Error500{Message: message}
}

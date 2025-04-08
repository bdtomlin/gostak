package handler

import (
	"net/http"

	"github.com/bdtomlin/gostak/internal/pages"
	"github.com/bdtomlin/gostak/internal/repo"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, _ := repo.ListAuthors()

	pages.AuthorList(authors).Render(r.Context(), w)
}

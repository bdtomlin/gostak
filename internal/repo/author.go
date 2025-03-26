package repo

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/model"
)

func InsertAuthor(a *model.Author) (*model.Author, error) {
	res, err := DB.Exec(
		"INSERT INTO authors (name, email) VALUES ($1, $2)",
		a.Name,
		a.Email,
	)
	if err != nil {
		return a, err
	}
	fmt.Println(res)
	return a, nil
}

func ListAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := DB.Select(&authors, "select * from authors")
	if err != nil {
		return authors, err
	}
	return authors, nil
}

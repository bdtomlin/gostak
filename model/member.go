package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Member struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	JoinDate string `db:"join_date"`
}

func (m Member) insert(db *sqlx.DB) (Member, error) {
	res, err := db.Exec(
		"INSERT INTO authors (name, email) VALUES ($1, $2)",
		"Bryan Tomlin",
		"bryan@tomlin.email",
	)
	if err != nil {
		return m, err
	}
	fmt.Println(res)
	return m, nil
}

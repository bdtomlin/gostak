package model

type Author struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

package model

import "database/sql"

type Book struct {
	ID            int            `db:"id"`
	Title         string         `db:"title"`
	AuthorID      int            `db:"author_id"`
	PublishedYear int            `db:"published_year"`
	Genre         sql.NullString `db:"genre"`
}

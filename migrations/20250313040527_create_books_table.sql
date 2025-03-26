-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
	id bigserial PRIMARY KEY,
	title TEXT NOT NULL,
	author_id INTEGER,
	published_year INTEGER,
	genre TEXT,
	FOREIGN KEY(author_id) REFERENCES authors(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table books;
-- +goose StatementEnd



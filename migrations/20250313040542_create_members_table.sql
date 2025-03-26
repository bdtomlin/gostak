-- +goose Up
-- +goose StatementBegin
CREATE TABLE members (
	id bigserial PRIMARY KEY,
	name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
	join_date TEXT NOT NULL DEFAULT CURRENT_DATE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table members;
-- +goose StatementEnd



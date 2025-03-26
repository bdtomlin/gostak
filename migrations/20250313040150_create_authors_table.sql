-- +goose Up
-- +goose StatementBegin
create table authors (
	id bigserial PRIMARY KEY,
    name text not null,
    email text unique not null
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table authors;
-- +goose StatementEnd



-- +goose Up
-- +goose StatementBegin
create or replace function update_updated_at()
returns trigger
as $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
language plpgsql
;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop function if exists update_updated_at()
;
-- +goose StatementEnd



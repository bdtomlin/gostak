-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT uuidv7(),
  inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null,
  first_name text not null,
  last_name text not null,
  email text unique not null,
  hashed_password text 
);
CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_updated_at();
create index idx_users_inserted_at on users (inserted_at);
create index idx_users_first_name on users (first_name);
create index idx_users_last_name on users (last_name);
create index idx_users_email on users (email);
create index idx_users_hashed_password on users (hashed_password);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trigger_update_updated_at ON users;
drop table users;
-- +goose StatementEnd



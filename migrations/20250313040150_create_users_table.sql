-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT uuidv7(),
  inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null,
  first_name text not null,
  last_name text not null,
  email text unique not null
);
CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_updated_at();
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trigger_update_updated_at
ON users;
drop table users;
-- +goose StatementEnd



-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
	id UUID PRIMARY KEY DEFAULT uuidv7(),
	user_id UUID REFERENCES users(id),
  token_hash TEXT UNIQUE NOT NULL,
  inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd



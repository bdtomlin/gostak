package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         string    `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	TokenHash  string    `db:"token_hash"`
	InsertedAt time.Time `db:"inserted_at"`
	Token      string
}

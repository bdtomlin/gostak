package model

import "time"

type Session struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	TokenHash  string    `db:"token_hash"`
	InsertedAt time.Time `db:"inserted_at"`
}

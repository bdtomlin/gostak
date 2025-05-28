package model

import (
	"github.com/google/uuid"
)

type SessionRepo struct{}

func (sr SessionRepo) CreateSession(userID uuid.UUID) (*Session, error) {
	// 1. create session token
	return nil, nil
}

func (sr SessionRepo) GetSessionUser(userID uuid.UUID, token string) (*User, error) {
	return nil, nil
}

package model

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/bdtomlin/gostak/internal/rand"
	"github.com/google/uuid"
)

const (
	MinBytesPerToken = 32
)

type SessionRepo struct {
	BytesPerToken int
}

func (sr SessionRepo) CreateSession(userID uuid.UUID) (*Session, error) {
	if sr.BytesPerToken < MinBytesPerToken {
		sr.BytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(sr.BytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("CreateSession: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: sr.hash(token),
	}
	query := `INSERT INTO sessions (user_id, token_hash) 
						VALUES (:user_id, :token_hash) 
            RETURNING *`
	_, err = DB.NamedExec(query, &session)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateSession: %w", err)
	}

	return &session, nil
}

func (sr SessionRepo) GetSessionUser(userID uuid.UUID, token string) (*User, error) {
	return nil, nil
}

func (ss *SessionRepo) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

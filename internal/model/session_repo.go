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

func (sr SessionRepo) GetSessionUser(userID *uuid.UUID, token string) (*User, error) {
	tokenHash := sr.hash(token)
	user := User{}
	err := DB.Get(&user, `
		select users.* from users
		inner join sessions on users.id = sessions.user_id
		where sessions.user_id = $1
		and sessions.token_hash = $2
		`, userID.String(), tokenHash)
	if err != nil {
		return nil, fmt.Errorf("repo.GetSessionUser: %w", err)
	}

	return &user, nil
}

func (sr SessionRepo) DeleteSession(userID *uuid.UUID, token string) error {
	tokenHash := sr.hash(token)
	_, err := DB.Exec(`
		delete from sessions
		where user_id = $1
		and token_hash = $2
		`, userID.String(), tokenHash)
	if err != nil {
		return fmt.Errorf("DeleteSession: %w", err)
	}

	return nil
}

func (sr SessionRepo) DeleteAllUserSessions(userID *uuid.UUID) error {
	_, err := DB.Exec(`
		delete from sessions
		where user_id = $1
		`, userID.String())
	if err != nil {
		return fmt.Errorf("DeleteAllUserSession: %w", err)
	}

	return nil
}

func (ss *SessionRepo) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

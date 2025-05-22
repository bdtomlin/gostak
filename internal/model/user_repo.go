package model

import (
	"fmt"

	"github.com/google/uuid"
)

type UserRepo struct{}

func (ur *UserRepo) ListUsers() ([]User, error) {
	var users []User
	err := DB.Select(&users, "select * from users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepo) UserEmailExists(email string) (bool, error) {
	var exists bool
	err := DB.Get(&exists, "select exists(select 1 from users where email = $1)", email)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *UserRepo) GetUser(strID string) (User, error) {
	var user User

	uuid, err := uuid.Parse(strID)
	if err != nil {
		return user, fmt.Errorf("repo.GetUser: %w", err)
	}

	err = DB.Get(&user, "select * from users where id = $1", uuid)
	if err != nil {
		return user, fmt.Errorf("repo.GetUser: %w", err)
	}
	return user, nil
}

func (ur *UserRepo) InsertUser(user User) error {
	query := `INSERT INTO users (email, first_name, last_name, hashed_password) 
						VALUES (:email, :first_name, :last_name, :hashed_password) 
            RETURNING *`
	_, err := DB.NamedExec(query, &user)
	if err != nil {
		return fmt.Errorf("repo.InsertUser: %w", err)
	}
	return nil
}

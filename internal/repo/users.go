package repo

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/google/uuid"
)

func ListUsers() ([]model.User, error) {
	var users []model.User
	err := DB.Select(&users, "select * from users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUser(strID string) (model.User, error) {
	var user model.User

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

package repo

import (
	"github.com/bdtomlin/gostak/internal/model"
)

func ListUsers() ([]model.User, error) {
	var users []model.User
	err := DB.Select(&users, "select * from users")
	if err != nil {
		return users, err
	}
	return users, nil
}

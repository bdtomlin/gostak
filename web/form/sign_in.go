package form

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/model"
)

type SignIn struct {
	*Form          `schema:"-"`
	Email          string
	Password       string
	HashedPassword string `schema:"-"`
	UserRepo       model.UserRepo
}

func NewSignIn() *SignIn {
	return &SignIn{
		Form: NewForm(),
	}
}

func (si *SignIn) Submit() (*model.User, error) {
	if !si.IsValid() {
		return nil, fmt.Errorf("Invalid Sign Up")
	}
	user, err := si.UserRepo.AuthenticateUser(si.Email, si.Password)

	if err != nil {
		si.AddError("Form", "Invalid email/password combination")
		return nil, err
	}

	return user, nil
}

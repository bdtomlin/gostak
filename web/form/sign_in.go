package form

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/bdtomlin/gostak/internal/model"
)

type SignIn struct {
	*Form          `schema:"-"`
	Email          string
	Password       string
	HashedPassword string `schema:"-"`
	UserRepo       model.UserRepo
	SessionRepo    model.SessionRepo
}

func NewSignIn() *SignIn {
	return &SignIn{
		Form: NewForm(),
	}
}

func (si *SignIn) Submit() (*model.Session, error) {
	user, err := si.UserRepo.AuthenticateUser(si.Email, si.Password)

	if err != nil {
		si.Password = ""
		return nil, errors.New("Invalid email/password combination")
	}

	session, err := si.SessionRepo.CreateSession(user.ID)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("form.SignUp.Submit: %w", err)
	}

	return session, nil
}

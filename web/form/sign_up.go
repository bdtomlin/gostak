package form

import (
	"fmt"
	"log/slog"
	"net/mail"
	"strings"

	"github.com/bdtomlin/gostak/internal/auth"
	"github.com/bdtomlin/gostak/internal/model"
)

type SignUp struct {
	*Form          `schema:"-"`
	FirstName      string
	LastName       string
	Email          string
	Password       string
	HashedPassword string            `schema:"-"`
	UserRepo       model.UserRepo    `schema:"-"`
	SessionRepo    model.SessionRepo `schema:"-"`
}

func NewSignUp() *SignUp {
	su := SignUp{
		Form: NewForm(),
	}
	su.AddValidators(
		su.ValidateEmail,
		su.ValidatePassword,
		su.ValidateFirstName,
		su.ValidateLastName,
	)

	return &su
}

func (su *SignUp) ValidateEmail() {
	address, err := mail.ParseAddress(su.Email)
	if err != nil {
		su.AddError("Email", "is invalid")
		return
	} else {
		su.Email = address.Address
	}
	exists, err := su.UserRepo.UserEmailExists(su.Email)
	if err != nil {
		su.AddError("Email", err.Error())
		return
	}
	if exists {
		su.AddError("Email", "email already exists")
	}
}

func (su *SignUp) ValidatePassword() {
	if strings.TrimSpace(su.Password) == "" {
		su.AddError("Password", "is required")
		return
	}
	err := auth.ValidatePassword(su.Password)
	if err != nil {
		su.AddError("Password", err.Error())
	}
	hp, err := auth.HashPassword(su.Password)
	if err != nil {
		su.AddError("Password", err.Error())
	} else {
		su.HashedPassword = hp
	}
}

func (su *SignUp) ValidateFirstName() {
	if strings.TrimSpace(su.FirstName) == "" {
		su.AddError("FirstName", "is required")
		return
	}
}

func (su *SignUp) ValidateLastName() {
	if strings.TrimSpace(su.LastName) == "" {
		su.AddError("LastName", "is required")
		return
	}
}

func (su *SignUp) Submit() (*model.Session, error) {
	if !su.IsValid() {
		return nil, fmt.Errorf("Invalid Sign Up")
	}
	user := model.User{
		Email:          su.Email,
		FirstName:      su.FirstName,
		LastName:       su.LastName,
		HashedPassword: su.HashedPassword,
	}

	err := su.UserRepo.InsertUser(&user)

	if err != nil {
		err = fmt.Errorf("form.SignUp.Submit: %w", err)
		slog.Error(err.Error())
		return nil, err
	}

	session, err := su.SessionRepo.CreateSession(user.ID)
	if err != nil {
		err = fmt.Errorf("form.SignUp.Submit: %w", err)
		slog.Error(err.Error())
		return nil, err
	}

	return session, nil
}

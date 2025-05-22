package form

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/bdtomlin/gostak/internal/model"
)

type SignUp struct {
	*Form     `schema:"-"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	UserRepo  model.UserRepo
}

func NewSignUp() *SignUp {
	su := &SignUp{
		Form: NewForm(),
	}
	su.AddValidator(su.ValidateEmail)
	su.AddValidator(su.ValidatePassword)
	su.AddValidator(su.ValidateFirstName)
	su.AddValidator(su.ValidateLastName)

	return su
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
	if len(su.Password) < 5 {
		su.AddError("Password", "is too short")
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

func (su *SignUp) Submit() error {
	if !su.IsValid() {
		return fmt.Errorf("Invalid Sign Up")
	}
	err := su.UserRepo.InsertUser(model.User{
		Email:          su.Email,
		FirstName:      su.FirstName,
		LastName:       su.LastName,
		HashedPassword: su.Password,
	})

	if err != nil {
		err = fmt.Errorf("form.SignUp.Submit: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

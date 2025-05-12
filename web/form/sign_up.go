package form

import (
	"net/mail"
	"strings"
)

type SignUp struct {
	*Form    `schema:"-"`
	Email    string
	Password string
}

func NewSignUp() *SignUp {
	su := &SignUp{
		Form: NewForm(),
	}
	su.AddValidator(su.ValidateEmail)
	su.AddValidator(su.ValidatePassword)

	return su
}

func (su *SignUp) ValidateEmail() {
	address, err := mail.ParseAddress(su.Email)
	if err != nil {
		su.AddError("Email", "is invalid")
	} else {
		su.Email = address.Address
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

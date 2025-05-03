package form

import (
	"strings"
)

type SignUp struct {
	*Form    `schema:"-"`
	Email    string
	Password string
}

func NewSignUp() *SignUp {
	su := SignUp{
		Form: NewForm(),
	}
	su.AddValidator(su.ValidateEmail)

	return &su
}

func (su *SignUp) ValidateEmail() {
	if !strings.Contains(su.Email, "@") {
		su.AddError("Email", "is invalid")
	}
}

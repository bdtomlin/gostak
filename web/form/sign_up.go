package form

import (
	"net/mail"
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

package form

import (
	"fmt"
	"strings"
)

type Form struct {
	Errors     map[string][]string
	ErrorCount int
	Validators []func()
}

func NewForm() *Form {
	return &Form{
		Errors:     map[string][]string{},
		Validators: []func(){},
	}
}

func (f *Form) AddValidator(validator func()) {
	f.AddValidators(validator)
}

func (f *Form) AddValidators(validators ...func()) {
	f.Validators = append(f.Validators, validators...)
}

func (f *Form) IsValid() bool {
	f.Validate()
	if len(f.Errors) > 0 {
		return false
	}
	return true
}

func (f *Form) Validate() {
	f.ResetValidation()
	for _, fn := range f.Validators {
		fn()
	}
}

func (f *Form) ResetValidation() {
	f.ErrorCount = 0
	f.Errors = map[string][]string{}
}

func (f *Form) ErrorsAny(field string) bool {
	if _, ok := f.Errors[field]; ok {
		return true
	}
	return false
}

func (f *Form) ErrorsOn(field string) []string {
	if f.ErrorsAny(field) {
		return f.Errors[field]
	}
	return []string{}
}

func (f *Form) InputErrors(field string) string {
	errsWithName := []string{}
	for _, err := range f.ErrorsOn(field) {
		errsWithName = append(errsWithName, fmt.Sprintf("%s %s", field, err))
	}
	return strings.Join(errsWithName, ", ")
}

func (f *Form) AddError(k string, v string) {
	f.ErrorCount++
	if _, ok := f.Errors[k]; ok {
		f.Errors[k] = append(f.Errors[k], v)
	} else {
		f.Errors[k] = []string{v}
	}
}

package models

import "github.com/sashamihalache/meetmeup/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)
	v.EqualToField("password", r.Password, "password_confirm", r.ConfirmPassword)

	v.Required("username", r.Username)
	v.MinLength("username", r.Username, 2)

	return v.IsValid(), v.Errors
}
func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}

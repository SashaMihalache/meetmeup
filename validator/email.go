package validator

import (
	"fmt"
	"regexp"
)

var emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !emailRegexp.MatchString(email) {
		v.Errors[field] = fmt.Sprintf("%s is not a valid email address", field)
		return false
	}

	return true
}

package validator

import "fmt"

func (v *Validator) MinLength(field, value string, minChar int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < minChar {
		v.Errors[field] = fmt.Sprintf("%s must be at least %d characters long", field, minChar)
		return false
	}

	return true
}

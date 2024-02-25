package validator

import "fmt"

// used interface{} to accept any type of value
func (v *Validator) EqualToField(field string, value interface{}, toEqualField string, toEquaValue interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if value != toEquaValue {
		v.Errors[field] = fmt.Sprintf("%s does not match %s", field, toEqualField)
		return false
	}

	return true
}

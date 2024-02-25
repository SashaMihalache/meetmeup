package validator

import "reflect"

func (v *Validator) Required(field, value string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if IsEmpty(value) {
		v.Errors[field] = field + " is required"
		return false
	}

	return true
}

func IsEmpty(value interface{}) bool {
	t := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return t.Len() == 0

	}

	return false
}

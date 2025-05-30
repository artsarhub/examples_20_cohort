package functions

import "errors"

func ApplyOperation(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func ValidateLogin(v string) error {
	if v == "" {
		return errors.New("invalid login")
	}
	return nil
}

func ValidatePassword(v string) error {
	if len(v) <= 8 {
		return errors.New("invalid password")
	}
	return nil
}

func Validate(v string, validator func(string) error) error {
	return validator(v)
}

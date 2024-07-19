package validation

import (
	"errors"
	"net/mail"
)

func ValidateEmail(email string, required bool) error {
    if email == "" && !required {
        return nil
    }

    if email == "" && required {
        return errors.New("Email is required")
    }

    _, err := mail.ParseAddress(email)
    if err!=nil {
        return errors.New("Invalid email")
    }
	return nil
}

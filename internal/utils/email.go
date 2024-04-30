package utils

import (
	"errors"
	"net/mail"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("El correo es obligatorio")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("El correo no es valido")
	}

	return nil
}

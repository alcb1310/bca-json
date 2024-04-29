package utils

import "errors"

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("La contrasena es obligatoria")
	}

	return nil
}

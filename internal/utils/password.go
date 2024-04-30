package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("La contraseña es obligatoria")
	}

	return nil
}

func EncryptPasssword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), 8)
}

func ComparePassword(hashed, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return errors.New("Credenciales inválidas")
	}

	return nil
}

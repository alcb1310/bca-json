package utils_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/utils"
)

func TestValidatePassword(t *testing.T) {
	t.Run("should return error if password is empty", func(t *testing.T) {
		password := ""
		expectedError := errors.New("La contraseña es obligatoria")

		err := utils.ValidatePassword(password)

		assert.Equal(t, expectedError, err)
	})

	t.Run("should return nil if password is not empty", func(t *testing.T) {
		password := "password"
		var expectedError error = nil

		err := utils.ValidatePassword(password)

		assert.Equal(t, expectedError, err)
	})
}

func TestEncryptPasssword(t *testing.T) {
	t.Run("should encrypt password", func(t *testing.T) {
		password := "password"

		_, err := utils.EncryptPasssword(password)

		assert.Nil(t, err)
	})

	t.Run("should decrypt password", func(t *testing.T) {
		password := "password"

		pass, err := utils.EncryptPasssword(password)

		assert.Nil(t, err)

		err = utils.ComparePassword(string(pass), password)

		assert.Nil(t, err)
	})

	t.Run("should return error if incorrect password", func(t *testing.T) {
		password := "password"
		expectedError := errors.New("Credenciales inválidas")

		pass, err := utils.EncryptPasssword(password)

		assert.Nil(t, err)

		err = utils.ComparePassword(string(pass), "incorrect")

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

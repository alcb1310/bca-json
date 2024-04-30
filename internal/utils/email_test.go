package utils_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/utils"
)

func TestValidateEmail(t *testing.T) {
	t.Run("should return error if email is empty", func(t *testing.T) {
		email := ""
		expectedError := errors.New("El correo es obligatorio")

		err := utils.ValidateEmail(email)

		assert.Equal(t, expectedError, err)
	})

	t.Run("should return nil if email is not empty", func(t *testing.T) {
		email := "a@a.com"
		var expectedError error = nil

		err := utils.ValidateEmail(email)

		assert.Equal(t, expectedError, err)
	})

	t.Run("should return error if email is not valid", func(t *testing.T) {
		email := "a"
		expectedError := errors.New("El correo no es valido")

		err := utils.ValidateEmail(email)

		assert.Equal(t, expectedError, err)
	})
}

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
		expectedError := errors.New("La contrasena es obligatoria")

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

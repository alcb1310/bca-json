package validation_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/validation"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		testName  string
		name      string
		minLength int
		required  bool
		expected  error
	}{
		{
			testName:  "Password is required",
			name:      "",
			minLength: 3,
			required:  true,
			expected:  errors.New("Password is required"),
		},
		{
			testName:  "Password is not required",
			name:      "",
			minLength: 3,
			required:  false,
			expected:  nil,
		},
		{
			testName:  "Password is invalid",
			name:      "a",
			minLength: 3,
			required:  true,
			expected:  errors.New("Invalid password"),
		},
		{
			testName:  "Password is valid",
			name:      "abc",
			minLength: 3,
			required:  true,
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result := validation.ValidatePassword(tt.name, tt.minLength, tt.required)
			assert.Equal(t, tt.expected, result)
		})
	}
}

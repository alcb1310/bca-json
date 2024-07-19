package validation_test

import (
	"errors"
	"testing"

	"github.com/alcb1310/bca-json/internals/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name     string
        email    string
        required bool
        expected error
    }{
        {
            name:     "Email is required",
            email:    "",
            required: true,
            expected: errors.New("Email is required"),
        },
        {
            name:     "Email is not required",
            email:    "",
            required: false,
            expected: nil,
        },
        {
            name:     "Email is invalid",
            email:    "a",
            required: true,
            expected: errors.New("Invalid email"),
        },
        {
            name:     "Email is valid",
            email:    "a@a.com",
            required: true,
            expected: nil,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            actual := validation.ValidateEmail(test.email, test.required)
            assert.Equal(t, test.expected, actual)
        })
    }
}

package validation_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/validation"
)

func TestValidateRuc(t *testing.T) {
	tests := []struct {
		name     string
		ruc      string
		required bool
		expected error
	}{
		{
			name:     "Invalid Length",
			ruc:      "1234",
			required: true,
			expected: errors.New("Invalid ID Length"),
		},
		{
			name:     "Invalid Length",
			ruc:      "12345678900",
			required: true,
			expected: errors.New("Invalid ID Length"),
		},
		{
			name:     "Invalid Length",
			ruc:      "12345678900000000",
			required: true,
			expected: errors.New("Invalid ID Length"),
		},
		{
			name:     "RUC is required",
			ruc:      "",
			required: true,
			expected: errors.New("ID is required"),
		},
		{
			name:     "Empty RUC when it is not required",
			ruc:      "",
			required: false,
			expected: nil,
		},
		{
			name:     "Cédula inválida",
			ruc:      "1234567890",
			required: true,
			expected: errors.New("Invalid ID"),
		},
		{
			name:     "Provincia inválida",
			ruc:      "3993366553001",
			required: false,
			expected: errors.New("Invalid ID"),
		},
		{
			name:     "RUC inválido",
			ruc:      "1934567890002",
			required: true,
			expected: errors.New("Invalid ID"),
		},
		{
			name:     "RUC válido",
			ruc:      "1934567890001",
			required: true,
			expected: nil,
		},
		{
			name:     "RUC válido",
			ruc:      "0195088608001",
			required: true,
			expected: nil,
		},
		{
			name:     "Cédula válida",
			ruc:      "1704749652",
			required: true,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateRuc(tt.ruc, tt.required)
			assert.Equal(t, tt.expected, err)
		})
	}
}

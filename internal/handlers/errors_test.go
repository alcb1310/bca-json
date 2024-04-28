package handlers_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/server"
	"github.com/alcb1310/bca-json/internal/utils_test"
)

func TestNotFound(t *testing.T) {
	s := server.NewServer()
	s.MountHandlers()

	t.Run("foo", func(t *testing.T) {
		expected := make(map[string]interface{})
		expected["error"] = "No se encuentra la ruta"
		expectedByte, _ := json.Marshal(expected)
		expectedString := string(expectedByte)

		request := httptest.NewRequest("GET", "/invalid", nil)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		assert.Equal(t, expectedString, strings.Trim(response.Body.String(), "\n"))
	})
}

func TestMethodNotAllowed(t *testing.T) {
	s := server.NewServer()
	s.MountHandlers()

	t.Run("foo", func(t *testing.T) {
		expected := make(map[string]interface{})
		expected["error"] = "Método no permitido"
		expectedByte, _ := json.Marshal(expected)
		expectedString := string(expectedByte)

		request := httptest.NewRequest("POST", "/foo", nil)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 405, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		assert.Equal(t, expectedString, strings.Trim(response.Body.String(), "\n"))
	})
}

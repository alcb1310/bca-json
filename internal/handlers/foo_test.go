package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/server"
	utils_test "github.com/alcb1310/bca-json/internal/utils"
)

func TestFoo(t *testing.T) {
	s := server.NewServer()
	s.MountHandlers()

	t.Run("foo", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/foo", nil)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Hello, World!", response.Body.String())
	})
}

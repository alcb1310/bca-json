package server_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/mocks"
	"github.com/alcb1310/bca-json/internal/server"
	"github.com/alcb1310/bca-json/internal/utils_test"
)

func TestUserLogin(t *testing.T) {
	t.Run("Should return 200 on POST /login", func(t *testing.T) {
		db := mocks.NewDatabaseMock()
		s := server.NewServer(db)
		s.MountHandlers()

		credentials := make(map[string]interface{})
		credentials["email"] = "test@test.com"
		credentials["password"] = "password123"

		if err := json.NewEncoder(&buf).Encode(credentials); err != nil {
			t.Fatal(err)
		}

		request := httptest.NewRequest("POST", "/login", &buf)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	})

	t.Run("Validate JSON information", func(t *testing.T) {

		t.Run("Empty JSON", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			request := httptest.NewRequest("POST", "/login", nil)
			response := utils_test.ExecuteRequest(request, s)

			errorResponse := make(map[string]interface{})
			errorResponse["error"] = "No se recibió JSON"

			errorString, _ := json.Marshal(errorResponse)

			assert.Equal(t, 400, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
			assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
		})

		t.Run("Validate Email", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty Email", func(t *testing.T) {
				credentials := make(map[string]interface{})

				if err := json.NewEncoder(&buf).Encode(credentials); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/login", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "email"
				error["message"] = "El correo es obligatorio"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})

			t.Run("Invalid Email", func(t *testing.T) {
				credentials := make(map[string]interface{})
				credentials["email"] = "invalid"

				if err := json.NewEncoder(&buf).Encode(credentials); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/login", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "email"
				error["message"] = "El correo no es valido"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})

		t.Run("Validate Password", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty Password", func(t *testing.T) {
				credentials := make(map[string]interface{})
				credentials["email"] = "test@test.com"

				if err := json.NewEncoder(&buf).Encode(credentials); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/login", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "password"
				error["message"] = "La contraseña es obligatoria"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})
	})
}

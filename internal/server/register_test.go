package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internal/mocks"
	"github.com/alcb1310/bca-json/internal/server"
	"github.com/alcb1310/bca-json/internal/types"
	"github.com/alcb1310/bca-json/internal/utils_test"
)

var buf bytes.Buffer
var emp uint = 1

func TestRegister(t *testing.T) {
	t.Run("Should return 201 on POST /register", func(t *testing.T) {
		db := mocks.NewDatabaseMock()
		s := server.NewServer(db)
		s.MountHandlers()

		registerData := make(map[string]interface{})
		registerData["ruc"] = "123456789"
		registerData["name"] = "test"
		registerData["employees"] = emp
		registerData["email"] = "test@test.com"
		registerData["password"] = "password123"
		registerData["username"] = "test"

		if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
			t.Fatal(err)
		}

		reg := types.RegisterInformation{
			RUC:       registerData["ruc"].(string),
			Name:      registerData["name"].(string),
			Email:     registerData["email"].(string),
			Password:  registerData["password"].(string),
			UserName:  registerData["username"].(string),
			Employees: &emp,
		}
		db.On("Register", reg).Return(uuid.New(), nil)
		request := httptest.NewRequest("POST", "/register", &buf)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	})

	t.Run("Validate JSON information", func(t *testing.T) {
		t.Run("Empty JSON", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			request := httptest.NewRequest("POST", "/register", &buf)
			response := utils_test.ExecuteRequest(request, s)

			errorResponse := make(map[string]interface{})
			errorResponse["error"] = "No se recibió JSON"

			errorString, _ := json.Marshal(errorResponse)

			assert.Equal(t, 400, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
			assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
		})

		t.Run("Validate RUC", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty RUC", func(t *testing.T) {
				registerData := make(map[string]interface{})
				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}
				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)
				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "ruc"
				error["message"] = "Ingrese un RUC"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})

		t.Run("Validate Name", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty Name", func(t *testing.T) {
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "nombre"
				error["message"] = "Ingrese un Nombre"
				errorResponse["error"] = error
				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})

		t.Run("Validate Email", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty Email", func(t *testing.T) {
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
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
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"
				registerData["email"] = "invalid"

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
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
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"
				registerData["email"] = "valid@mail.com"

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
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

		t.Run("Validate UserName", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Empty Password", func(t *testing.T) {
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"
				registerData["email"] = "valid@mail.com"
				registerData["password"] = "password1234"

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "username"
				error["message"] = "El nombre de usuario es obligatorio"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})

		t.Run("Validate Employees", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			t.Run("Negative number", func(t *testing.T) {
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"
				registerData["email"] = "valid@mail.com"
				registerData["password"] = "password1234"
				registerData["username"] = "test user"
				registerData["employees"] = -1

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "employees"
				error["message"] = "Los empleados deben ser un número positivo válido"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})

			t.Run("Not a number", func(t *testing.T) {
				registerData := make(map[string]interface{})
				registerData["ruc"] = "123456789"
				registerData["name"] = "test"
				registerData["email"] = "valid@mail.com"
				registerData["password"] = "password1234"
				registerData["username"] = "test user"
				registerData["employees"] = "test"

				if err := json.NewEncoder(&buf).Encode(registerData); err != nil {
					t.Fatal(err)
				}

				request := httptest.NewRequest("POST", "/register", &buf)
				response := utils_test.ExecuteRequest(request, s)

				errorResponse := make(map[string]interface{})
				errorResponse["error"] = make(map[string]interface{})
				error := make(map[string]interface{})
				error["field"] = "employees"
				error["message"] = "Los empleados deben ser un número positivo válido"
				errorResponse["error"] = error

				errorString, _ := json.Marshal(errorResponse)

				assert.Equal(t, 400, response.Code)
				assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
				assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
			})
		})
	})

	t.Run("Conflict on POST /register", func(t *testing.T) {
		regData := types.RegisterInformation{
			RUC:       "123456789",
			Name:      "test",
			Email:     "valid@mail",
			Password:  "password1234",
			UserName:  "test user",
			Employees: &emp,
		}

		t.Run("Duplicate RUC", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			if err := json.NewEncoder(&buf).Encode(regData); err != nil {
				t.Fatal(err)
			}
			db.On("Register", regData).Return(uuid.Nil, errors.New("violates unique constraint \"company_ruc_unique\""))
			request := httptest.NewRequest("POST", "/register", &buf)
			response := utils_test.ExecuteRequest(request, s)

			errorResponse := make(map[string]interface{})
			errorResponse["error"] = make(map[string]interface{})
			error := make(map[string]interface{})
			error["field"] = "ruc"
			error["message"] = "Ya existe una empresa con este ruc"
			errorResponse["error"] = error

			errorString, _ := json.Marshal(errorResponse)

			assert.Equal(t, 409, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
			assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
		})

		t.Run("Duplicate Name", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			if err := json.NewEncoder(&buf).Encode(regData); err != nil {
				t.Fatal(err)
			}
			db.On("Register", regData).Return(uuid.Nil, errors.New("violates unique constraint \"company_name_unique\""))
			request := httptest.NewRequest("POST", "/register", &buf)
			response := utils_test.ExecuteRequest(request, s)

			errorResponse := make(map[string]interface{})
			errorResponse["error"] = make(map[string]interface{})
			error := make(map[string]interface{})
			error["field"] = "name"
			error["message"] = "Ya existe una empresa con este nombre"
			errorResponse["error"] = error

			errorString, _ := json.Marshal(errorResponse)

			assert.Equal(t, 409, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
			assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
		})

		t.Run("Duplicate Email", func(t *testing.T) {
			db := mocks.NewDatabaseMock()
			s := server.NewServer(db)
			s.MountHandlers()

			if err := json.NewEncoder(&buf).Encode(regData); err != nil {
				t.Fatal(err)
			}

			db.On("Register", regData).Return(uuid.Nil, errors.New("violates unique constraint \"user_email_unique\""))
			request := httptest.NewRequest("POST", "/register", &buf)
			response := utils_test.ExecuteRequest(request, s)

			errorResponse := make(map[string]interface{})
			errorResponse["error"] = make(map[string]interface{})
			error := make(map[string]interface{})
			error["field"] = "email"
			error["message"] = "Ya existe un usuario con este correo"
			errorResponse["error"] = error

			errorString, _ := json.Marshal(errorResponse)

			assert.Equal(t, 409, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
			assert.Equal(t, string(errorString), strings.Trim(response.Body.String(), "\n"))
		})
	})

	t.Run("Unknown error on POST /register", func(t *testing.T) {
		regData := types.RegisterInformation{
			RUC:       "123456789",
			Name:      "test",
			Email:     "valid@mail",
			Password:  "password1234",
			UserName:  "test user",
			Employees: &emp,
		}
		db := mocks.NewDatabaseMock()
		s := server.NewServer(db)
		s.MountHandlers()

		if err := json.NewEncoder(&buf).Encode(regData); err != nil {
			t.Fatal(err)
		}

		db.On("Register", regData).Return(uuid.Nil, errors.New("unknown error"))

		request := httptest.NewRequest("POST", "/register", &buf)
		response := utils_test.ExecuteRequest(request, s)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		assert.Equal(t, "\"unknown error\"", strings.Trim(response.Body.String(), "\n"))
	})
}

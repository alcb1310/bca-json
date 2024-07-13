package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/mocks"
)

func TestCreateCompany(t *testing.T) {
	t.Run("Data validation", func(t *testing.T) {
		t.Run("Empty body", func(t *testing.T) {
			s := server.New(mocks.NewService(t))
			if s == nil {
				t.Fatal("Server should not be nil")
			}

			req, err := http.NewRequest("POST", "/api/v2/companies", nil)
			if err != nil {
				t.Fatal(err)
			}

			expectedResponse := map[string]string{
				"error": "Missing Body",
			}
			var actualResponse map[string]string

			response := getResponse(t, s, req, http.StatusBadRequest)
			err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, expectedResponse, actualResponse)
		})

		t.Run("Should have company RUC", func(t *testing.T) {
			t.Run("Empty RUC", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}

				expectedResponse := map[string]string{
					"error": "ID is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid RUC", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1234567890001"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Invalid ID",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Should have company name", func(t *testing.T) {
			t.Run("Empty name", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Name is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid name", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "a"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Invalid name",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Should have a user email", func(t *testing.T) {
			t.Run("Empty email", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Email is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid email", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"
				data["email"] = "a"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Invalid email",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Should have the user's name", func(t *testing.T) {
			t.Run("empty name", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"
				data["email"] = "a@b.c"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Name is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
			t.Run("invalid name", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"
				data["email"] = "a@b.c"
				data["user_name"] = "a"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Invalid name",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})
		t.Run("Should have the user's password", func(t *testing.T) {
			t.Run("empty password", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"
				data["email"] = "a@b.c"
				data["user_name"] = "alc"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Password is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("invalid password", func(t *testing.T) {
				s := server.New(mocks.NewService(t))
				if s == nil {
					t.Fatal("Server should not be nil")
				}

				data := make(map[string]string)
				data["ruc"] = "1791838300001"
				data["name"] = "Company Name"
				data["email"] = "a@b.c"
				data["user_name"] = "alc"
				data["password"] = "a"

				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse := map[string]string{
					"error": "Invalid password",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})
	})
}

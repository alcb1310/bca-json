package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/mocks"
)

var token string

func TestCreateUser(t *testing.T) {
	// TODO: implement
	token = loginUser(t)

	t.Run("Data validation", func(t *testing.T) {
		t.Run("No body sent", func(t *testing.T) {
			s := server.New(mocks.NewService(t), "secret")
			assert.NotNil(t, s)

			req, err := http.NewRequest("POST", "/api/v2/bca/users", nil)
			assert.NoError(t, err)

			expectedResponse := map[string]string{
				"error": "Missing Body",
			}
			var actualResponse map[string]string

			response := getResponse(t, s, req, http.StatusBadRequest, token)
			err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
			assert.NoError(t, err)
			assert.Equal(t, expectedResponse, actualResponse)
		})

		t.Run("Email", func(t *testing.T) {
			t.Run("No email", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Email is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid email", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["email"] = "test"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid email",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Password", func(t *testing.T) {
			t.Run("No password", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["email"] = "test@test.com"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Password is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid password", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["email"] = "test@test.com"
				data["password"] = "te"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid password",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Name", func(t *testing.T) {
			t.Run("No name", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["email"] = "test@test.com"
				data["password"] = "test"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Name is required",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Invalid name", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["email"] = "test@test.com"
				data["password"] = "test"
				data["name"] = "a"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid name",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})
	})

	t.Run("Valid data", func(t *testing.T) {
		user := types.CreateUser{
			User: types.User{
				Name:      "Test User",
				Email:     "test@test.com",
				RoleID:    "a",
				CompanyID: companyID,
			},
			Password: "test",
		}
		dataBase := mocks.NewService(t)
		s := server.New(dataBase, "secret")
		assert.NotNil(t, s)

		data := make(map[string]string)
		data["email"] = "test@test.com"
		data["password"] = "test"
		data["name"] = "Test User"

		dataBase.EXPECT().
			CreateUser(user).
			Return(
				types.User{
					ID:        uuid.UUID{},
					Name:      "Test User",
					Email:     "test@test.com",
					RoleID:    "a",
					CompanyID: companyID,
				},
				nil,
			).Times(1)
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(data)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/api/v2/bca/users", &buf)
		assert.NoError(t, err)

		response := getResponse(t, s, req, http.StatusCreated, token)

		expectedResponse := map[string]interface{}{
			"user": types.User{
				ID:        uuid.UUID{},
				Name:      "Test User",
				Email:     "test@test.com",
				RoleID:    "a",
				CompanyID: companyID,
            },
		}
		var actualResponse map[string]types.User
		err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse["user"], actualResponse["user"])
	})
}

func TestUpdateUser(t *testing.T) {
    t.Run("Invalid data", func(t *testing.T) {
        t.Run("Email", func(t *testing.T) {
            uuid := uuid.New()
            s := server.New(mocks.NewService(t), "secret")
            assert.NotNil(t, s)

            data := make(map[string]string)
            data["email"] = "a"
            var buf bytes.Buffer
            err := json.NewEncoder(&buf).Encode(data)
            assert.NoError(t, err)

            req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v2/bca/users/%s", uuid.String()), &buf)
            assert.NoError(t, err)

            expectedResponse := map[string]string{
                "error": "Invalid email",
            }
            var actualResponse map[string]string

            response := getResponse(t, s, req, http.StatusBadRequest, token)
            err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
            assert.NoError(t, err)
            assert.Equal(t, expectedResponse, actualResponse)
        })

        t.Run("Name", func(t *testing.T) {
            uuid := uuid.New()
            s := server.New(mocks.NewService(t), "secret")
            assert.NotNil(t, s)

            data := make(map[string]string)
            data["name"] = "a"
            var buf bytes.Buffer
            err := json.NewEncoder(&buf).Encode(data)
            assert.NoError(t, err)

            req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v2/bca/users/%s", uuid.String()), &buf)
            assert.NoError(t, err)

            expectedResponse := map[string]string{
                "error": "Invalid name",
            }
            var actualResponse map[string]string

            response := getResponse(t, s, req, http.StatusBadRequest, token)
            err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
            assert.NoError(t, err)
            assert.Equal(t, expectedResponse, actualResponse)
        })

        t.Run("Password", func(t *testing.T) {
            uuid := uuid.New()
            s := server.New(mocks.NewService(t), "secret")
            assert.NotNil(t, s)

            data := make(map[string]string)
            data["password"] = "a"
            var buf bytes.Buffer
            err := json.NewEncoder(&buf).Encode(data)
            assert.NoError(t, err)

            req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v2/bca/users/%s", uuid.String()), &buf)
            assert.NoError(t, err)

            expectedResponse := map[string]string{
                "error": "Invalid password",
            }
            var actualResponse map[string]string

            response := getResponse(t, s, req, http.StatusBadRequest, token)
            err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
            assert.NoError(t, err)
            assert.Equal(t, expectedResponse, actualResponse)
        })
    })
}

package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/mocks"
)

func TestCreateProject(t *testing.T) {
	token = loginUser(t)

	t.Run("Data validation", func(t *testing.T) {
		t.Run("No body sent", func(t *testing.T) {
			s := server.New(mocks.NewService(t), "secret")
			assert.NotNil(t, s)

			req, err := http.NewRequest("POST", "/api/v2/bca/projects", nil)
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

		t.Run("Empty body sent", func(t *testing.T) {
			s := server.New(mocks.NewService(t), "secret")
			assert.NotNil(t, s)

			data := make(map[string]string)
			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(data)
			assert.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
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

		t.Run("Invalid gross area", func(t *testing.T) {
			t.Run("Non numeric gross area", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]string)
				data["name"] = "abc"
				data["gross_area"] = "a"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid gross area",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Negative gross area", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]interface{})
				data["name"] = "abc"
				data["gross_area"] = -1.0
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid gross area",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})
		})

		t.Run("Invalid net area", func(t *testing.T) {
			t.Run("Non numeric net area", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]interface{})
				data["name"] = "abc"
				data["gross_area"] = 1
				data["net_area"] = "a"
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid net area",
				}
				var actualResponse map[string]string

				response := getResponse(t, s, req, http.StatusBadRequest, token)
				err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, actualResponse)
			})

			t.Run("Negative net area", func(t *testing.T) {
				s := server.New(mocks.NewService(t), "secret")
				assert.NotNil(t, s)

				data := make(map[string]interface{})
				data["name"] = "abc"
				data["gross_area"] = 1
				data["net_area"] = -1.0
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(data)
				assert.NoError(t, err)

				req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
				assert.NoError(t, err)

				expectedResponse := map[string]string{
					"error": "Invalid net area",
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
		project := types.Project{
			Name:      "abc",
			GrossArea: 1.0,
			NetArea:   1.0,
            IsActive: true,
			CompanyID: companyID,
		}
		dataBase := mocks.NewService(t)
		s := server.New(dataBase, "secret")
		assert.NotNil(t, s)

		data := make(map[string]interface{})
		data["name"] = "abc"
		data["gross_area"] = 1
		data["net_area"] = 1
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(data)
		assert.NoError(t, err)

		dataBase.EXPECT().CreateProject(project).Return(
			types.Project{
				ID:          uuid.UUID{},
				Name:        "abc",
				GrossArea:   1.0,
				NetArea:     1.0,
				CompanyID:   companyID,
				IsActive:    true,
				LastClosure: nil,
			},
			nil,
		).Times(1)

		req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
		assert.NoError(t, err)

		expectedResponse := map[string]types.Project{
			"project": {
				ID:        uuid.UUID{},
				Name:      "abc",
				GrossArea: 1.0,
				NetArea:   1.0,
				CompanyID: companyID,
				IsActive:  true,
			},
		}
		var actualResponse map[string]types.Project

		response := getResponse(t, s, req, http.StatusCreated, token)
		err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse["project"], actualResponse["project"])
	})
}

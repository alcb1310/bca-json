package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/mocks"
)

var companyID = uuid.MustParse("4b114c26-b038-4cfa-ae6e-ad46c73ef59d")

func getResponse(t *testing.T, s *server.Server, r *http.Request, code int, token string) *httptest.ResponseRecorder {
	r.Header.Set("Content-Type", "application/json")
	if token != "" {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	w := httptest.NewRecorder()
	s.Server.ServeHTTP(w, r)

	assert.Equal(t, code, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	return w
}

func loginUser(t *testing.T) string {
	dataBase := mocks.NewService(t)
	s := server.New(dataBase, "secret")
	assert.NotNil(t, s)

	dataBase.EXPECT().
		Login("test@test.com", "test").
		Return(
			types.User{
				ID:        uuid.New(),
				Name:      "Test User",
				Email:     "test@test.com",
				RoleID:    "a",
				CompanyID: companyID,
			},
			nil,
		)

	req, err := http.NewRequest(
		"POST",
		"/api/v2/login",
		bytes.NewBuffer([]byte(`{"email": "test@test.com", "password": "test"}`)),
	)
	assert.NoError(t, err)

	response := getResponse(t, s, req, http.StatusOK, "")
	var actualResponse map[string]string
	err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	return actualResponse["token"]
}

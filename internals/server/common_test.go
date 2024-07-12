package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcb1310/bca-json/internals/server"
)

func getResponse(t *testing.T, s *server.Server, r *http.Request, code int) *httptest.ResponseRecorder {
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.Server.ServeHTTP(w, r)

	assert.Equal(t, code, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	return w
}

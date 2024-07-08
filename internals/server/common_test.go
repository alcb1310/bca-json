package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/stretchr/testify/assert"
)

func getResponse(t *testing.T, s *server.Server, r *http.Request, code int) *httptest.ResponseRecorder {
    w := httptest.NewRecorder()
    s.Server.ServeHTTP(w, r)

    assert.Equal(t, code, w.Code)
    assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

    return w
}

package utils_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/alcb1310/bca-json/internal/server"
)

// ExecuteRequest executes a request and returns the response recorder
func ExecuteRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	s.Router.ServeHTTP(response, req)

	return response
}

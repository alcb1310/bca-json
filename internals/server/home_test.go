package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
    s := server.New()

    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    response := httptest.NewRecorder()
    s.Server.ServeHTTP(response, req)

    assert.Equal(t, http.StatusOK, response.Code)

    expectedResponse := map[string]string{
        "message": "Hello World!",
    }

    var actualResponse map[string]string
    err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
    if err != nil {
        t.Fatal(err)
    }

    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, expectedResponse, actualResponse)
    assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

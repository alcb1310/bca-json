package server_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
    db := mocks.NewService(t)
    s := server.New(db)

    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    response := getResponse(t, s, req, http.StatusOK)

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
}

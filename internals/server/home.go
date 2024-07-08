package server

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) error {
    res := map[string]string{
        "message": "Hello World!",
    }

    err := json.NewEncoder(w).Encode(res)
	return err
}

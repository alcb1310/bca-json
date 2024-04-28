package handlers

import (
	"fmt"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

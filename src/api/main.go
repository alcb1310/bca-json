package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/alcb1310/bca-json/internals/server"
)


func main() {
    s := server.New()

    port := "8081"
    slog.Info("Starting server", "port", port)

    if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Server); err != nil {
        panic(fmt.Sprintf("Failed to start server: %v", err))
    }
}

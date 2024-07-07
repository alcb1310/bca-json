package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/alcb1310/bca-json/internals/server"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
    s := server.New()

    port := os.Getenv("PORT")
    slog.Info("Starting server", "port", port)

    if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Server); err != nil {
        panic(fmt.Sprintf("Failed to start server: %v", err))
    }
}

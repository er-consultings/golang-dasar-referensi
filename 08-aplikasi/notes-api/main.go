// Command notes-api runs a simple notes API at http://localhost:8090.
//
//	go run ./08-aplikasi/notes-api
//	curl -X POST localhost:8090/notes -d '{"title":"Learning Go","body":"slices chapter"}'
//	curl localhost:8090/notes
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"golang-dasar-referensi/08-aplikasi/notes-api/internal/notes"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func run() error {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8090"
	}
	srv := &http.Server{
		Addr:              addr,
		Handler:           notes.Handler(notes.NewStore()),
		ReadHeaderTimeout: 5 * time.Second,
	}
	slog.Info("server running", "addr", addr)
	return srv.ListenAndServe()
}

// Command notes-api runs a simple notes API at http://localhost:8090.
//
//	go run ./08-aplikasi/notes-api
//	curl -X POST localhost:8090/notes -d '{"title":"Learning Go","body":"slices chapter"}'
//	curl localhost:8090/notes
//
// Press Ctrl+C to stop: requests that are still running get up to 5 seconds to finish.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// ctx is cancelled on Ctrl+C (SIGINT) or when a process manager sends SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errc := make(chan error, 1)
	go func() {
		slog.Info("server running", "addr", addr)
		errc <- srv.ListenAndServe()
	}()

	select {
	case err := <-errc: // the server could not start, for example the port is taken
		return err
	case <-ctx.Done():
	}

	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}
	if err := <-errc; !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	slog.Info("server stopped cleanly")
	return nil
}

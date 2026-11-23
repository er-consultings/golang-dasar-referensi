// Command catatan-api menjalankan API catatan sederhana di http://localhost:8090.
//
//	go run ./08-aplikasi/catatan-api
//	curl -X POST localhost:8090/catatan -d '{"judul":"Belajar Go","isi":"bab slice"}'
//	curl localhost:8090/catatan
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"golang-dasar-referensi/08-aplikasi/catatan-api/internal/catatan"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server berhenti", "err", err)
		os.Exit(1)
	}
}

func run() error {
	alamat := os.Getenv("ALAMAT")
	if alamat == "" {
		alamat = ":8090"
	}
	srv := &http.Server{
		Addr:              alamat,
		Handler:           catatan.Handler(catatan.NewPenyimpan()),
		ReadHeaderTimeout: 5 * time.Second,
	}
	slog.Info("server berjalan", "alamat", alamat)
	return srv.ListenAndServe()
}

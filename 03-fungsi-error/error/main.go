// Error sebagai nilai: sentinel, membungkus dengan %w, errors.Is, errors.As.
//
//	go run ./03-fungsi-error/error
package main

import (
	"errors"
	"fmt"
)

// ErrTidakDitemukan adalah error sentinel yang bisa dikenali pemanggil.
var ErrTidakDitemukan = errors.New("data tidak ditemukan")

// ErrValidasi membawa detail kolom yang salah.
type ErrValidasi struct {
	Field string
}

func (e *ErrValidasi) Error() string { return "isian tidak valid: " + e.Field }

type Produk struct {
	ID   int
	Nama string
}

var gudang = map[int]Produk{1: {ID: 1, Nama: "Kopi"}}

func cari(id int) (Produk, error) {
	p, ok := gudang[id]
	if !ok {
		return Produk{}, ErrTidakDitemukan
	}
	return p, nil
}

func muatProduk(id int) (Produk, error) {
	if id <= 0 {
		return Produk{}, &ErrValidasi{Field: "id"}
	}
	p, err := cari(id)
	if err != nil {
		return Produk{}, fmt.Errorf("memuat produk %d: %w", id, err)
	}
	return p, nil
}

func main() {
	for _, id := range []int{1, 7, -1} {
		p, err := muatProduk(id)
		var ev *ErrValidasi
		switch {
		case err == nil:
			fmt.Println("ketemu:", p.Nama)
		case errors.Is(err, ErrTidakDitemukan):
			fmt.Println("404:", err)
		case errors.As(err, &ev):
			fmt.Println("400, periksa kolom", ev.Field)
		default:
			fmt.Println("error lain:", err)
		}
	}
}

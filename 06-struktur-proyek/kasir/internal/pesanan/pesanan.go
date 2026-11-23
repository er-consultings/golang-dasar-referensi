// Package pesanan membuat pesanan dan menghitung totalnya.
package pesanan

import (
	"errors"
	"fmt"

	"golang-dasar-referensi/06-struktur-proyek/kasir/internal/produk"
)

// ErrPesananKosong dikembalikan bila pesanan tidak berisi barang.
var ErrPesananKosong = errors.New("pesanan kosong")

// Katalog adalah satu-satunya hal yang dibutuhkan paket ini dari katalog produk.
// Interface didefinisikan di sisi pemakai, jadi test bisa memberi versi palsu.
type Katalog interface {
	Cari(id string) (produk.Produk, error)
}

// Baris adalah permintaan satu produk dengan jumlahnya.
type Baris struct {
	ProdukID string
	Jumlah   int
}

// BarisPesanan adalah satu baris yang sudah diberi nama dan harga.
type BarisPesanan struct {
	Nama   string
	Harga  int64
	Jumlah int
}

// Subtotal mengembalikan harga dikali jumlah.
func (b BarisPesanan) Subtotal() int64 { return b.Harga * int64(b.Jumlah) }

// Pesanan berisi baris-baris yang sudah dihargai.
type Pesanan struct {
	Baris []BarisPesanan
}

// Total menjumlahkan semua subtotal.
func (p Pesanan) Total() int64 {
	var t int64
	for _, b := range p.Baris {
		t += b.Subtotal()
	}
	return t
}

// Layanan membuat pesanan memakai sebuah Katalog.
type Layanan struct {
	katalog Katalog
}

// NewLayanan membuat Layanan.
func NewLayanan(k Katalog) *Layanan { return &Layanan{katalog: k} }

// Buat memeriksa setiap baris dan mengisi nama serta harga dari katalog.
func (l *Layanan) Buat(baris []Baris) (Pesanan, error) {
	if len(baris) == 0 {
		return Pesanan{}, ErrPesananKosong
	}
	out := Pesanan{Baris: make([]BarisPesanan, 0, len(baris))}
	for _, b := range baris {
		if b.Jumlah <= 0 {
			return Pesanan{}, fmt.Errorf("jumlah %s harus lebih dari nol", b.ProdukID)
		}
		p, err := l.katalog.Cari(b.ProdukID)
		if err != nil {
			return Pesanan{}, fmt.Errorf("produk %s: %w", b.ProdukID, err)
		}
		out.Baris = append(out.Baris, BarisPesanan{Nama: p.Nama, Harga: p.Harga, Jumlah: b.Jumlah})
	}
	return out, nil
}

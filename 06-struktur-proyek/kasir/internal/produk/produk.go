// Package produk menyimpan katalog produk yang dijual.
package produk

import "errors"

// ErrTidakDitemukan dikembalikan bila ID produk tidak ada di katalog.
var ErrTidakDitemukan = errors.New("produk tidak ditemukan")

// Produk adalah barang yang dijual. Harga dalam rupiah.
type Produk struct {
	ID    string
	Nama  string
	Harga int64
}

// Katalog menyimpan produk di memori.
type Katalog struct {
	byID map[string]Produk
}

// NewKatalog membuat katalog dari daftar produk.
func NewKatalog(daftar []Produk) *Katalog {
	k := &Katalog{byID: make(map[string]Produk, len(daftar))}
	for _, p := range daftar {
		k.byID[p.ID] = p
	}
	return k
}

// Cari mengembalikan produk dengan ID tersebut.
func (k *Katalog) Cari(id string) (Produk, error) {
	p, ok := k.byID[id]
	if !ok {
		return Produk{}, ErrTidakDitemukan
	}
	return p, nil
}

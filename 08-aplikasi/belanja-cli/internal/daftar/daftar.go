// Package daftar menyimpan barang belanjaan dan menghitung totalnya.
package daftar

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ErrJumlahTidakValid dikembalikan bila jumlah atau harga tidak positif.
var ErrJumlahTidakValid = errors.New("jumlah dan harga harus lebih dari nol")

// ErrTidakAda dikembalikan saat menghapus barang yang tidak ada.
var ErrTidakAda = errors.New("barang tidak ada di daftar")

// Barang adalah satu baris di daftar belanja. Harga dalam rupiah.
type Barang struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Harga  int64  `json:"harga"`
}

// Subtotal mengembalikan harga dikali jumlah.
func (b Barang) Subtotal() int64 { return b.Harga * int64(b.Jumlah) }

// Daftar adalah kumpulan barang. Zero value-nya langsung bisa dipakai.
type Daftar struct {
	Barang []Barang `json:"barang"`
}

// Tambah memasukkan barang baru, atau menambah jumlah bila namanya sudah ada.
func (d *Daftar) Tambah(nama string, jumlah int, harga int64) error {
	if jumlah <= 0 || harga <= 0 {
		return ErrJumlahTidakValid
	}
	for i := range d.Barang {
		if d.Barang[i].Nama == nama {
			d.Barang[i].Jumlah += jumlah
			d.Barang[i].Harga = harga
			return nil
		}
	}
	d.Barang = append(d.Barang, Barang{Nama: nama, Jumlah: jumlah, Harga: harga})
	return nil
}

// Hapus membuang barang dengan nama tersebut.
func (d *Daftar) Hapus(nama string) error {
	for i := range d.Barang {
		if d.Barang[i].Nama == nama {
			d.Barang = append(d.Barang[:i], d.Barang[i+1:]...)
			return nil
		}
	}
	return ErrTidakAda
}

// Total menjumlahkan subtotal semua barang.
func (d *Daftar) Total() int64 {
	var total int64
	for _, b := range d.Barang {
		total += b.Subtotal()
	}
	return total
}

// Muat membaca daftar dari berkas JSON. Berkas yang belum ada menghasilkan daftar kosong.
func Muat(path string) (*Daftar, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return &Daftar{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("membaca %s: %w", path, err)
	}
	var d Daftar
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("format %s rusak: %w", path, err)
	}
	return &d, nil
}

// Simpan menulis daftar ke berkas JSON.
func (d *Daftar) Simpan(path string) error {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

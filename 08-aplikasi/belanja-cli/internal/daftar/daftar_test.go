package daftar

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestTambahDanTotal(t *testing.T) {
	var d Daftar // zero value langsung dipakai
	if err := d.Tambah("kopi", 2, 18000); err != nil {
		t.Fatal(err)
	}
	if err := d.Tambah("kopi", 1, 18000); err != nil {
		t.Fatal(err)
	}
	if len(d.Barang) != 1 {
		t.Fatalf("barang = %d baris, ingin 1", len(d.Barang))
	}
	if got, want := d.Total(), int64(54000); got != want {
		t.Errorf("Total() = %d, ingin %d", got, want)
	}
}

func TestTambahTidakValid(t *testing.T) {
	var d Daftar
	if err := d.Tambah("roti", 0, 15000); !errors.Is(err, ErrJumlahTidakValid) {
		t.Errorf("err = %v, ingin ErrJumlahTidakValid", err)
	}
}

func TestHapus(t *testing.T) {
	var d Daftar
	_ = d.Tambah("kopi", 1, 18000)
	_ = d.Tambah("roti", 1, 15000)
	if err := d.Hapus("kopi"); err != nil {
		t.Fatal(err)
	}
	if len(d.Barang) != 1 || d.Barang[0].Nama != "roti" {
		t.Errorf("sisa = %+v", d.Barang)
	}
	if err := d.Hapus("teh"); !errors.Is(err, ErrTidakAda) {
		t.Errorf("err = %v, ingin ErrTidakAda", err)
	}
}

func TestSimpanLaluMuat(t *testing.T) {
	path := filepath.Join(t.TempDir(), "belanja.json") // folder sementara, dihapus otomatis
	var d Daftar
	_ = d.Tambah("kopi", 2, 18000)
	if err := d.Simpan(path); err != nil {
		t.Fatal(err)
	}
	got, err := Muat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.Total() != 36000 {
		t.Errorf("Total setelah dimuat = %d", got.Total())
	}
}

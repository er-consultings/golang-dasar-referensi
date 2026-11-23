package produk

import (
	"errors"
	"testing"
)

func TestCari(t *testing.T) {
	k := NewKatalog([]Produk{{ID: "kopi", Nama: "Kopi", Harga: 18000}})
	p, err := k.Cari("kopi")
	if err != nil || p.Harga != 18000 {
		t.Fatalf("Cari(kopi) = %+v, %v", p, err)
	}
	if _, err := k.Cari("teh"); !errors.Is(err, ErrTidakDitemukan) {
		t.Errorf("Cari(teh) err = %v, ingin ErrTidakDitemukan", err)
	}
}

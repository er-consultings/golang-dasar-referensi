package pesanan

import (
	"errors"
	"testing"

	"golang-dasar-referensi/06-struktur-proyek/kasir/internal/produk"
)

// katalogPalsu memenuhi interface Katalog tanpa perlu paket produk yang asli.
type katalogPalsu map[string]produk.Produk

func (k katalogPalsu) Cari(id string) (produk.Produk, error) {
	p, ok := k[id]
	if !ok {
		return produk.Produk{}, produk.ErrTidakDitemukan
	}
	return p, nil
}

func TestBuat(t *testing.T) {
	l := NewLayanan(katalogPalsu{"kopi": {ID: "kopi", Nama: "Kopi", Harga: 18000}})

	tests := []struct {
		name      string
		baris     []Baris
		wantTotal int64
		wantErr   error
	}{
		{"satu produk", []Baris{{"kopi", 2}}, 36000, nil},
		{"kosong", nil, 0, ErrPesananKosong},
		{"tidak ada di katalog", []Baris{{"teh", 1}}, 0, produk.ErrTidakDitemukan},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := l.Buat(tt.baris)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("err = %v, ingin %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := p.Total(); got != tt.wantTotal {
				t.Errorf("Total() = %d, ingin %d", got, tt.wantTotal)
			}
		})
	}
}

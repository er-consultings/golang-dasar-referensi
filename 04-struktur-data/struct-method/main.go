// Struct, konstruktor, method dengan value dan pointer receiver, tag JSON.
//
//	go run ./04-struktur-data/struct-method
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int64  `json:"harga"`
	Stok  int    `json:"stok"`
	modal int64  // huruf kecil: tidak ikut ke JSON
}

func NewProduk(nama string, harga int64) (*Produk, error) {
	if harga < 0 {
		return nil, errors.New("harga tidak boleh negatif")
	}
	return &Produk{Nama: nama, Harga: harga}, nil
}

func (p *Produk) Total(jumlah int) int64 { return p.Harga * int64(jumlah) }

func (p *Produk) KurangiStok(n int) error {
	if n > p.Stok {
		return fmt.Errorf("stok %s tinggal %d", p.Nama, p.Stok)
	}
	p.Stok -= n
	return nil
}

func main() {
	p, err := NewProduk("Kopi Arabika", 85000)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Stok = 3
	p.modal = 60000
	fmt.Printf("%+v\n", *p)
	fmt.Println("total 2:", p.Total(2))

	if err := p.KurangiStok(5); err != nil {
		fmt.Println("gagal:", err)
	}
	_ = p.KurangiStok(2)
	fmt.Println("sisa stok:", p.Stok)

	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	var dariJSON Produk
	if err := json.Unmarshal([]byte(`{"id":2,"nama":"Teh","harga":12000}`), &dariJSON); err == nil {
		fmt.Printf("%+v\n", dariJSON)
	}
}

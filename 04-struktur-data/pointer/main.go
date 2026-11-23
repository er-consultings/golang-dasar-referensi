// Pointer: & dan *, mengubah nilai pemanggil, nil pointer.
//
//	go run ./04-struktur-data/pointer
package main

import "fmt"

type Produk struct {
	Nama  string
	Harga int64
}

func diskon(p *Produk, persen int64) {
	p.Harga -= p.Harga * persen / 100
}

func buatProduk() *Produk {
	p := Produk{Nama: "Baru"}
	return &p // aman, kompiler memindahkannya ke heap
}

func main() {
	x := 10
	p := &x
	*p = 20
	fmt.Println("x:", x)

	kopi := Produk{Nama: "Kopi", Harga: 50000}
	diskon(&kopi, 10)
	fmt.Println("harga setelah diskon:", kopi.Harga)

	fmt.Println(buatProduk().Nama)

	var kosong *Produk
	if kosong == nil {
		fmt.Println("pointer nil, jangan akses field-nya")
	}
}

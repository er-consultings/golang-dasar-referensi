// Variabel, konstanta, zero value, shadowing, dan iota.
//
//	go run ./02-dasar-bahasa/variabel
package main

import "fmt"

// Status memakai iota sebagai enumerasi.
type Status int

const (
	StatusDraf Status = iota
	StatusTerbit
	StatusArsip
)

// String membuat Status tercetak sebagai teks, bukan angka.
func (s Status) String() string {
	switch s {
	case StatusDraf:
		return "draf"
	case StatusTerbit:
		return "terbit"
	case StatusArsip:
		return "arsip"
	}
	return fmt.Sprintf("Status(%d)", int(s))
}

// Variabel tingkat paket wajib memakai var.
var aplikasi = "kasir"

func main() {
	var kota string = "Bandung"
	var umur = 25
	nama := "Sari"
	fmt.Println(aplikasi, kota, umur, nama)

	// Zero value: setiap tipe punya nilai awal yang jelas.
	var (
		n     int
		f     float64
		s     string
		b     bool
		slice []int
		m     map[string]int
	)
	fmt.Printf("zero value: %d %v %q %v %v %v\n", n, f, s, b, slice == nil, m == nil)

	// Tukar nilai tanpa variabel bantu.
	a, c := 1, 2
	a, c = c, a
	fmt.Println("setelah ditukar:", a, c)

	// Shadowing: x di dalam blok adalah variabel baru.
	x := 1
	if true {
		x := 2
		fmt.Println("x di dalam blok:", x)
	}
	fmt.Println("x di luar blok:", x)

	fmt.Println("status:", StatusDraf, StatusTerbit, StatusArsip)
}

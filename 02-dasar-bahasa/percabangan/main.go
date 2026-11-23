// if, if dengan pernyataan awal, early return, dan switch.
//
//	go run ./02-dasar-bahasa/percabangan
package main

import (
	"fmt"
	"strconv"
)

// predikat memakai switch tanpa ekspresi sebagai pengganti rantai if else.
func predikat(nilai int) string {
	switch {
	case nilai >= 90:
		return "A"
	case nilai >= 80:
		return "B"
	case nilai >= 70:
		return "C"
	default:
		return "D"
	}
}

// diskon menangani kasus khusus lebih dulu (early return).
func diskon(umur int) float64 {
	if umur < 0 {
		return 0
	}
	if umur < 12 {
		return 0.5
	}
	return 0
}

func jenisHari(hari string) string {
	switch hari {
	case "Sabtu", "Minggu":
		return "akhir pekan"
	case "Jumat":
		return "hampir libur"
	default:
		return "hari kerja"
	}
}

func main() {
	for _, n := range []int{95, 81, 70, 40} {
		fmt.Println(n, "=>", predikat(n))
	}
	fmt.Println("diskon umur 8:", diskon(8))
	fmt.Println("Minggu:", jenisHari("Minggu"))

	for _, input := range []string{"17", "tujuh"} {
		if n, err := strconv.Atoi(input); err != nil {
			fmt.Println(input, "bukan angka")
		} else {
			fmt.Println(input, "genap?", n%2 == 0)
		}
	}
}

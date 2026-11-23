// Tipe data dasar, string, byte, rune, dan konversi.
//
//	go run ./02-dasar-bahasa/tipe-data
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	a, b := 7, 2
	fmt.Println("7 / 2 =", a/b)                           // 3
	fmt.Println("7 % 2 =", a%b)                           // 1
	fmt.Println("float:", float64(a)/float64(b))          // 3.5
	fmt.Println("0.1 + 0.2 =", 0.1+0.2 == 0.3, "(float)") // hati-hati dengan float

	// Uang disimpan sebagai bilangan bulat rupiah.
	var harga int64 = 18000
	fmt.Println("total 3 kopi:", harga*3)

	k := "Kopi ☕"
	fmt.Println("byte:", len(k), "karakter:", utf8.RuneCountInString(k))
	for i, r := range k {
		fmt.Printf("  indeks %d: %q\n", i, r)
	}

	// Konversi teks ke angka selalu bisa gagal, jadi periksa error.
	for _, input := range []string{"42", "42a"} {
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("gagal:", err)
			continue
		}
		fmt.Println("angka:", n)
	}
	fmt.Println("Itoa:", strconv.Itoa(65), "bukan", string(rune(65)))

	raw := `baris pertama
baris kedua \n tidak diproses`
	fmt.Println(raw)
}

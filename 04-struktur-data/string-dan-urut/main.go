// Alat yang sering dipakai di soal latihan: strings, strings.Builder,
// byte dan rune, unicode, serta mengurutkan dengan slices dan sort.
//
//	go run ./04-struktur-data/string-dan-urut
package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
)

type Siswa struct {
	Nama  string
	Nilai int
}

func main() {
	s := "  Halo, Dunia Go  "
	fmt.Printf("%q\n", strings.TrimSpace(s))
	fmt.Println(strings.ToLower(s), strings.ToUpper(s))
	fmt.Println(strings.Contains(s, "Go"), strings.Index(s, "Dunia"))
	fmt.Println(strings.Split("a,b,c", ","), strings.Fields(s))
	fmt.Println(strings.Join([]string{"x", "y"}, "-"), strings.Repeat("ab", 3))

	// Membangun string panjang: pakai strings.Builder, bukan + berulang.
	var b strings.Builder
	for i := range 5 {
		fmt.Fprintf(&b, "%d;", i)
	}
	fmt.Println(b.String())

	// String tidak bisa diubah. Ubah ke []byte atau []rune dulu.
	kata := []rune("kasur")
	slices.Reverse(kata)
	fmt.Println(string(kata))

	// Byte sebagai angka: menghitung huruf a sampai z.
	var hitung [26]int
	for _, c := range "banana" {
		hitung[c-'a']++
	}
	fmt.Println("a:", hitung['a'-'a'], "n:", hitung['n'-'a'])

	fmt.Println(unicode.IsLetter('x'), unicode.IsDigit('7'), string(unicode.ToLower('Q')))

	angka := []int{5, 2, 9, 1}
	slices.Sort(angka)
	fmt.Println(angka)

	siswa := []Siswa{{"Rina", 88}, {"Budi", 92}, {"Sari", 88}}
	// Urut nilai tertinggi, lalu nama A sampai Z.
	slices.SortFunc(siswa, func(a, b Siswa) int {
		if c := cmp.Compare(b.Nilai, a.Nilai); c != 0 {
			return c
		}
		return cmp.Compare(a.Nama, b.Nama)
	})
	fmt.Println(siswa)

	// Gaya lama yang masih sering terlihat di jawaban orang lain.
	sort.Slice(angka, func(i, j int) bool { return angka[i] > angka[j] })
	fmt.Println(angka)
}

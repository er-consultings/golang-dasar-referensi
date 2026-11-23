// Map: comma ok, menghitung frekuensi, map sebagai set, map nil.
//
//	go run ./04-struktur-data/map
package main

import (
	"fmt"
	"strings"
)

func main() {
	harga := map[string]int{"kopi": 18000, "teh": 12000}
	harga["roti"] = 15000
	delete(harga, "teh")

	if h, ok := harga["es"]; !ok {
		fmt.Println("es tidak tersedia, nilai zero:", h)
	}

	teks := "kopi susu kopi gula kopi"
	frekuensi := map[string]int{}
	for _, kata := range strings.Fields(teks) {
		frekuensi[kata]++
	}
	fmt.Println(frekuensi)

	sudahDilihat := map[int]struct{}{}
	for _, n := range []int{3, 1, 3, 2, 1} {
		if _, ada := sudahDilihat[n]; ada {
			fmt.Println("duplikat:", n)
			continue
		}
		sudahDilihat[n] = struct{}{}
	}

	type Koordinat struct{ X, Y int }
	peta := map[Koordinat]string{{0, 0}: "titik awal"}
	fmt.Println(peta[Koordinat{0, 0}])

	var m map[string]int
	fmt.Println("baca map nil aman:", m["a"])
	// m["a"] = 1 // panic: assignment to entry in nil map
}

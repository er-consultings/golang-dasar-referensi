// Slice: len, cap, append, potongan berbagi array, salinan, paket slices.
//
//	go run ./04-struktur-data/slice
package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []int
	for i := range 5 {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	nama := []string{"rina", "budi", "sari"}
	hasil := make([]string, 0, len(nama)) // kapasitas disiapkan di awal
	for _, n := range nama {
		hasil = append(hasil, "halo "+n)
	}
	fmt.Println(hasil)

	angka := []int{10, 20, 30, 40, 50}
	a := angka[1:3]
	a[0] = 99
	fmt.Println("potongan mengubah asli:", angka)

	salinan := slices.Clone(angka)
	salinan[0] = -1
	fmt.Println("salinan terpisah:", angka, salinan)

	fmt.Println("contains 30:", slices.Contains(angka, 30))
	fmt.Println("index 40:", slices.Index(angka, 40))
	slices.Sort(angka)
	fmt.Println("urut:", angka)
	i, ok := slices.BinarySearch(angka, 40)
	fmt.Println("binary search 40:", i, ok)

	// Slice 2 dimensi, sering muncul di soal matriks.
	grid := make([][]int, 3)
	for r := range grid {
		grid[r] = make([]int, 4)
	}
	grid[1][2] = 7
	fmt.Println(grid)

	var kosong []int
	fmt.Println("kosong?", len(kosong) == 0)
}

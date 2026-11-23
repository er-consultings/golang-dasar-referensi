// Semua bentuk for, range, salinan di range, urutan map, dan label.
//
//	go run ./02-dasar-bahasa/perulangan
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("(klasik)")

	saldo := 100
	for saldo > 0 {
		saldo -= 30
	}
	fmt.Println("saldo akhir:", saldo, "(seperti while)")

	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println("(range bilangan bulat, Go 1.22+)")

	angka := []int{1, 2, 3}
	for _, n := range angka {
		n *= 10 // hanya mengubah salinan
		_ = n
	}
	fmt.Println("tidak berubah:", angka)
	for i := range angka {
		angka[i] *= 10
	}
	fmt.Println("berubah:", angka)

	stok := map[string]int{"jeruk": 3, "apel": 5, "mangga": 1}
	for _, nama := range slices.Sorted(maps.Keys(stok)) {
		fmt.Println(nama, stok[nama])
	}

	tabel := [][]int{{1, 2}, {3, 4}, {5, 6}}
	target := 4
cari:
	for i, baris := range tabel {
		for j, sel := range baris {
			if sel == target {
				fmt.Printf("ketemu %d di [%d][%d]\n", target, i, j)
				break cari
			}
		}
	}
}

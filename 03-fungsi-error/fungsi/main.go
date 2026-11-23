// Fungsi: beberapa nilai kembali, variadic, fungsi sebagai nilai, closure.
//
//	go run ./03-fungsi-error/fungsi
package main

import (
	"errors"
	"fmt"
)

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}

func jumlah(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

func terapkan(angka []int, f func(int) int) []int {
	hasil := make([]int, 0, len(angka))
	for _, n := range angka {
		hasil = append(hasil, f(n))
	}
	return hasil
}

func penghitung() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func naikkan(n int)      { n++ }
func naikkanAsli(n *int) { *n++ }

func main() {
	if h, err := bagi(10, 4); err == nil {
		fmt.Println("10 / 4 =", h)
	}
	if _, err := bagi(1, 0); err != nil {
		fmt.Println("gagal:", err)
	}

	data := []int{4, 5, 6}
	fmt.Println(jumlah(), jumlah(1, 2, 3), jumlah(data...))

	kuadrat := func(n int) int { return n * n }
	fmt.Println(terapkan([]int{1, 2, 3}, kuadrat))

	hitung := penghitung()
	fmt.Println(hitung(), hitung(), hitung())

	skor := 5
	naikkan(skor)
	fmt.Println("setelah naikkan:", skor)
	naikkanAsli(&skor)
	fmt.Println("setelah naikkanAsli:", skor)
}

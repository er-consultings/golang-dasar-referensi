// Every form of for, range, copies in range, map order, and labels.
//
//	go run ./02-dasar-bahasa/loops
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
	fmt.Println("(classic)")

	balance := 100
	for balance > 0 {
		balance -= 30
	}
	fmt.Println("final balance:", balance, "(like while)")

	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println("(range over an integer, Go 1.22+)")

	numbers := []int{1, 2, 3}
	for _, n := range numbers {
		n *= 10 // only changes the copy
		_ = n
	}
	fmt.Println("unchanged:", numbers)
	for i := range numbers {
		numbers[i] *= 10
	}
	fmt.Println("changed:", numbers)

	stock := map[string]int{"orange": 3, "apple": 5, "mango": 1}
	for _, name := range slices.Sorted(maps.Keys(stock)) {
		fmt.Println(name, stock[name])
	}

	grid := [][]int{{1, 2}, {3, 4}, {5, 6}}
	target := 4
search:
	for i, row := range grid {
		for j, cell := range row {
			if cell == target {
				fmt.Printf("found %d at [%d][%d]\n", target, i, j)
				break search
			}
		}
	}
}

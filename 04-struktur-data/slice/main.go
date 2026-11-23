// Slices: len, cap, append, sub-slices sharing an array, copies, the slices package.
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

	names := []string{"rina", "budi", "sari"}
	greetings := make([]string, 0, len(names)) // capacity prepared up front
	for _, n := range names {
		greetings = append(greetings, "hello "+n)
	}
	fmt.Println(greetings)

	numbers := []int{10, 20, 30, 40, 50}
	a := numbers[1:3]
	a[0] = 99
	fmt.Println("sub-slice changes the original:", numbers)

	clone := slices.Clone(numbers)
	clone[0] = -1
	fmt.Println("separate copy:", numbers, clone)

	fmt.Println("contains 30:", slices.Contains(numbers, 30))
	fmt.Println("index of 40:", slices.Index(numbers, 40))
	slices.Sort(numbers)
	fmt.Println("sorted:", numbers)
	i, ok := slices.BinarySearch(numbers, 40)
	fmt.Println("binary search 40:", i, ok)

	// A 2D slice, common in matrix problems.
	grid := make([][]int, 3)
	for r := range grid {
		grid[r] = make([]int, 4)
	}
	grid[1][2] = 7
	fmt.Println(grid)

	var empty []int
	fmt.Println("empty?", len(empty) == 0)
}

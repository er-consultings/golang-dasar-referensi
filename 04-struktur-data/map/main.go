// Map: comma ok, counting frequencies, map as a set, nil map.
//
//	go run ./04-struktur-data/map
package main

import (
	"fmt"
	"strings"
)

func main() {
	prices := map[string]int{"coffee": 18000, "tea": 12000}
	prices["bread"] = 15000
	delete(prices, "tea")

	if p, ok := prices["ice"]; !ok {
		fmt.Println("ice is not available, zero value:", p)
	}

	text := "coffee milk coffee sugar coffee"
	frequency := map[string]int{}
	for _, word := range strings.Fields(text) {
		frequency[word]++
	}
	fmt.Println(frequency)

	seen := map[int]struct{}{}
	for _, n := range []int{3, 1, 3, 2, 1} {
		if _, ok := seen[n]; ok {
			fmt.Println("duplicate:", n)
			continue
		}
		seen[n] = struct{}{}
	}

	type Point struct{ X, Y int }
	places := map[Point]string{{0, 0}: "origin"}
	fmt.Println(places[Point{0, 0}])

	var m map[string]int
	fmt.Println("reading a nil map is safe:", m["a"])
	// m["a"] = 1 // panic: assignment to entry in nil map
}

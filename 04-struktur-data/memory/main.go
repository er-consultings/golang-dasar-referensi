// Stack, heap, escape analysis, and the garbage collector.
//
//	go run ./04-struktur-data/memory
//	go build -gcflags='-m -l' ./04-struktur-data/memory   # where does each value live?
//	go test -bench . -benchmem ./04-struktur-data/memory
package main

import (
	"fmt"
	"runtime"
	"slices"
)

type Point struct{ X, Y int }

// sum keeps p on the stack: p never leaves the function.
func sum(x, y int) int {
	p := Point{x, y}
	return p.X + p.Y
}

// newPoint returns a pointer to a local variable, so p "escapes" and moves to the heap.
func newPoint(x, y int) *Point {
	p := Point{x, y}
	return &p
}

// firstThree returns a small view of a big slice. The view keeps the whole
// big array alive, because both share the same underlying array.
func firstThree(big []int) []int {
	return big[:3]
}

// firstThreeCopy copies the three values, so the big array can be freed.
func firstThreeCopy(big []int) []int {
	return slices.Clone(big[:3])
}

func heapMB() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.HeapAlloc) / (1 << 20)
}

func main() {
	fmt.Println("sum:", sum(1, 2))
	p := newPoint(3, 4)
	fmt.Println("point:", *p)

	runtime.GC()
	fmt.Printf("heap at start:            %5.1f MB\n", heapMB())

	big := make([]int, 10_000_000) // about 76 MB
	big[0], big[1], big[2] = 7, 8, 9
	view := firstThree(big)
	big = nil
	runtime.GC()
	fmt.Printf("heap with a 3-item view:  %5.1f MB\n", heapMB())

	copied := firstThreeCopy(view)
	view = nil
	runtime.GC()
	fmt.Printf("heap with a 3-item copy:  %5.1f MB\n", heapMB())
	fmt.Println(copied)
}

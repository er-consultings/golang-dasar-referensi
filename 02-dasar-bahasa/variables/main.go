// Variables, constants, zero values, shadowing, and iota.
//
//	go run ./02-dasar-bahasa/variables
package main

import "fmt"

// Status uses iota as an enumeration.
type Status int

const (
	StatusDraft Status = iota
	StatusPublished
	StatusArchived
)

// String makes a Status print as text instead of a number.
func (s Status) String() string {
	switch s {
	case StatusDraft:
		return "draft"
	case StatusPublished:
		return "published"
	case StatusArchived:
		return "archived"
	}
	return fmt.Sprintf("Status(%d)", int(s))
}

// Package-level variables must use var.
var appName = "cashier"

func main() {
	var city string = "Bandung"
	var age = 25
	name := "Sari"
	fmt.Println(appName, city, age, name)

	// Zero value: every type has a well-defined starting value.
	var (
		n     int
		f     float64
		s     string
		b     bool
		slice []int
		m     map[string]int
	)
	fmt.Printf("zero values: %d %v %q %v %v %v\n", n, f, s, b, slice == nil, m == nil)

	// Swap values without a helper variable.
	a, c := 1, 2
	a, c = c, a
	fmt.Println("after swap:", a, c)

	// Shadowing: x inside the block is a new variable.
	x := 1
	if true {
		x := 2
		fmt.Println("x inside the block:", x)
	}
	fmt.Println("x outside the block:", x)

	fmt.Println("status:", StatusDraft, StatusPublished, StatusArchived)
}

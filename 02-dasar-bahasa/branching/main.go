// if, if with an init statement, early return, and switch.
//
//	go run ./02-dasar-bahasa/branching
package main

import (
	"fmt"
	"strconv"
)

// grade uses a switch without an expression instead of an if else chain.
func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "D"
	}
}

// discount handles the special cases first (early return).
func discount(age int) float64 {
	if age < 0 {
		return 0
	}
	if age < 12 {
		return 0.5
	}
	return 0
}

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Friday":
		return "almost weekend"
	default:
		return "weekday"
	}
}

func main() {
	for _, n := range []int{95, 81, 70, 40} {
		fmt.Println(n, "=>", grade(n))
	}
	fmt.Println("discount at age 8:", discount(8))
	fmt.Println("Sunday:", dayType("Sunday"))

	for _, input := range []string{"17", "seven"} {
		if n, err := strconv.Atoi(input); err != nil {
			fmt.Println(input, "is not a number")
		} else {
			fmt.Println(input, "even?", n%2 == 0)
		}
	}
}

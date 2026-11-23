// Functions: multiple return values, variadic, functions as values, closures.
//
//	go run ./03-fungsi-error/functions
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor must not be zero")
	}
	return a / b, nil
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func apply(numbers []int, f func(int) int) []int {
	result := make([]int, 0, len(numbers))
	for _, n := range numbers {
		result = append(result, f(n))
	}
	return result
}

func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func increment(n int)          { n++ }
func incrementOriginal(n *int) { *n++ }

func main() {
	if q, err := divide(10, 4); err == nil {
		fmt.Println("10 / 4 =", q)
	}
	if _, err := divide(1, 0); err != nil {
		fmt.Println("failed:", err)
	}

	data := []int{4, 5, 6}
	fmt.Println(sum(), sum(1, 2, 3), sum(data...))

	square := func(n int) int { return n * n }
	fmt.Println(apply([]int{1, 2, 3}, square))

	next := counter()
	fmt.Println(next(), next(), next())

	score := 5
	increment(score)
	fmt.Println("after increment:", score)
	incrementOriginal(&score)
	fmt.Println("after incrementOriginal:", score)
}

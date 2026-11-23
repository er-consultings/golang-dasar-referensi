// Package fizzbuzz: for every number from 1 to n, write "Fizz" if it is divisible by 3,
// "Buzz" if divisible by 5, "FizzBuzz" if divisible by both, otherwise the number itself.
//
// Similar to LeetCode 412. Pattern: loop and branching.
// Time O(n), space O(n) for the result.
package fizzbuzz

import "strconv"

func fizzBuzz(n int) []string {
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0: // check the most specific case first
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}

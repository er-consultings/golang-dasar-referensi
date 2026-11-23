// Package validparentheses: the string contains only ( ) [ ] { }. Return true if
// every opening bracket is closed by the same type and in the right order.
//
// Similar to LeetCode 20. Pattern: stack. In Go, a stack is just a slice with append and s[:len(s)-1].
// Time O(n), space O(n).
package validparentheses

func isValid(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, c := range s {
		open, isClosing := pairs[c]
		if !isClosing {
			stack = append(stack, c) // push
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1] // pop
	}
	return len(stack) == 0
}

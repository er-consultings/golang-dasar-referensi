// Package validparentheses: string hanya berisi ( ) [ ] { }. Kembalikan true bila
// setiap kurung buka ditutup oleh jenis yang sama dan dalam urutan yang benar.
//
// Mirip LeetCode 20. Pola: stack. Di Go, stack cukup slice dengan append dan s[:len(s)-1].
// Waktu O(n), ruang O(n).
package validparentheses

func isValid(s string) bool {
	pasangan := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, c := range s {
		buka, isTutup := pasangan[c]
		if !isTutup {
			stack = append(stack, c) // push
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != buka {
			return false
		}
		stack = stack[:len(stack)-1] // pop
	}
	return len(stack) == 0
}

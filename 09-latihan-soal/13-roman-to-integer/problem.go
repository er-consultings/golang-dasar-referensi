// Package romantointeger: convert a Roman numeral (I, V, X, L, C, D, M) to an integer.
// A smaller symbol before a larger one is subtracted, for example IV = 4.
//
// Similar to LeetCode 13. Pattern: a map of symbol values, compared with the next symbol.
// Time O(n), space O(1).
package romantointeger

var values = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := values[s[i]]
		if i+1 < len(s) && v < values[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

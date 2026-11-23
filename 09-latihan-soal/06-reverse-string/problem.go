// Package reversestring: reverse a byte slice in place, without creating a new slice.
//
// Similar to LeetCode 344. Pattern: two pointers swapping.
// Time O(n), space O(1).
package reversestring

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Package validanagram: two strings s and t contain lowercase letters a to z.
// Return true if t is a rearrangement of the letters of s.
//
// Similar to LeetCode 242. Pattern: count letter frequencies.
// Time O(n), space O(1) because the array always has 26 elements.
package validanagram

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++ // s[i] is a byte, 'a' is 97
		counts[t[i]-'a']--
	}
	for _, n := range counts {
		if n != 0 {
			return false
		}
	}
	return true
}

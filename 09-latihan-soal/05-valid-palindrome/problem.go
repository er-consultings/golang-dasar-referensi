// Package validpalindrome: after converting uppercase letters to lowercase and removing
// every character that is not a letter or a digit, does the string read the same
// forwards and backwards?
//
// Similar to LeetCode 125. Pattern: two pointers from both ends.
// Time O(n), space O(1).
package validpalindrome

func isAlnum(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func lower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlnum(s[i]) {
			i++
			continue
		}
		if !isAlnum(s[j]) {
			j--
			continue
		}
		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

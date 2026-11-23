// Package reversestring: balik urutan slice byte di tempat, tanpa membuat slice baru.
//
// Mirip LeetCode 344. Pola: dua pointer saling bertukar.
// Waktu O(n), ruang O(1).
package reversestring

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

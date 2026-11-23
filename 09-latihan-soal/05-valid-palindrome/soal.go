// Package validpalindrome: setelah huruf besar diubah ke kecil dan semua
// karakter selain huruf dan angka dibuang, apakah string terbaca sama dari
// depan dan belakang?
//
// Mirip LeetCode 125. Pola: dua pointer dari kedua ujung.
// Waktu O(n), ruang O(1).
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

// Package validanagram: dua string s dan t berisi huruf kecil a sampai z.
// Kembalikan true bila t adalah susunan ulang huruf-huruf s.
//
// Mirip LeetCode 242. Pola: hitung frekuensi huruf.
// Waktu O(n), ruang O(1) karena array selalu 26 elemen.
package validanagram

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var hitung [26]int
	for i := 0; i < len(s); i++ {
		hitung[s[i]-'a']++ // s[i] adalah byte, 'a' bernilai 97
		hitung[t[i]-'a']--
	}
	for _, n := range hitung {
		if n != 0 {
			return false
		}
	}
	return true
}

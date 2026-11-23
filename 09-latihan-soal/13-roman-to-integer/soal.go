// Package romantointeger: ubah angka Romawi (I, V, X, L, C, D, M) menjadi bilangan bulat.
// Simbol yang lebih kecil di depan simbol lebih besar berarti dikurangi, misalnya IV = 4.
//
// Mirip LeetCode 13. Pola: map nilai simbol, bandingkan dengan simbol berikutnya.
// Waktu O(n), ruang O(1).
package romantointeger

var nilai = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := nilai[s[i]]
		if i+1 < len(s) && v < nilai[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

// Package containsduplicate: kembalikan true bila ada angka yang muncul
// lebih dari sekali.
//
// Mirip LeetCode 217. Pola: map sebagai himpunan (set).
// Waktu O(n), ruang O(n).
package containsduplicate

func containsDuplicate(nums []int) bool {
	dilihat := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ada := dilihat[n]; ada {
			return true
		}
		dilihat[n] = struct{}{}
	}
	return false
}

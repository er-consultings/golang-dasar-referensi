// Package binarysearch: cari target di slice terurut naik. Kembalikan indeksnya,
// atau -1 bila tidak ada.
//
// Mirip LeetCode 704. Pola: binary search, separuh ruang pencarian dibuang setiap langkah.
// Waktu O(log n), ruang O(1).
package binarysearch

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // cara aman menghitung tengah tanpa overflow
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

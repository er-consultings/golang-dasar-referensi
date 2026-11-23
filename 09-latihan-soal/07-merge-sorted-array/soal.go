// Package mergesortedarray: nums1 berisi m angka terurut diikuti n slot kosong (0),
// nums2 berisi n angka terurut. Gabungkan nums2 ke nums1 sehingga nums1 terurut.
//
// Mirip LeetCode 88. Pola: isi dari belakang supaya tidak menimpa angka yang belum dipindah.
// Waktu O(m+n), ruang O(1).
package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// Package twosum: diberi slice angka dan target, kembalikan indeks dua angka
// yang jumlahnya sama dengan target. Dijamin ada tepat satu jawaban, dan
// elemen yang sama tidak boleh dipakai dua kali.
//
// Mirip LeetCode 1. Pola: hash map menyimpan "angka yang sudah dilihat".
package twosum

// twoSumBrute mencoba semua pasangan. Waktu O(n^2), ruang O(1).
// Cara pertama yang terpikir, benar tetapi lambat untuk n besar.
func twoSumBrute(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSum sekali jalan. Untuk setiap angka, pasangannya adalah target - angka.
// Kalau pasangan itu sudah pernah dilihat, jawaban ketemu.
// Waktu O(n), ruang O(n).
func twoSum(nums []int, target int) []int {
	dilihat := make(map[int]int, len(nums)) // angka -> indeks
	for i, n := range nums {
		if j, ok := dilihat[target-n]; ok {
			return []int{j, i}
		}
		dilihat[n] = i
	}
	return nil
}

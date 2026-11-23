// Package movezeroes: pindahkan semua 0 ke belakang tanpa mengubah urutan angka
// lain, di tempat.
//
// Mirip LeetCode 283. Pola: pointer tulis dan pointer baca.
// Waktu O(n), ruang O(1).
package movezeroes

func moveZeroes(nums []int) {
	tulis := 0
	for _, n := range nums {
		if n != 0 {
			nums[tulis] = n
			tulis++
		}
	}
	for i := tulis; i < len(nums); i++ {
		nums[i] = 0
	}
}

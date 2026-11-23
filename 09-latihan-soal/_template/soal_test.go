package template

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// Salin Example 1, 2, 3 dari soal, lalu tambah kasus tepi sendiri:
		// slice kosong, satu elemen, angka negatif, nilai terbesar.
		{[]int{1, 2, 3}, -1}, // ganti dengan Example 1 dari soal, test ini sengaja gagal dulu
	}
	for _, tt := range tests {
		if got := solve(tt.nums); got != tt.want {
			t.Errorf("solve(%v) = %d, ingin %d", tt.nums, got, tt.want)
		}
	}
}

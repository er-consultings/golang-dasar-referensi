// Package climbingstairs: tangga punya n anak tangga. Setiap langkah naik 1 atau 2.
// Ada berapa cara berbeda untuk sampai ke atas?
//
// Mirip LeetCode 70. Pola: cara(n) = cara(n-1) + cara(n-2), seperti Fibonacci.
package climbingstairs

// climbRekursif langsung dari rumusnya. Benar, tetapi waktu O(2^n)
// karena nilai yang sama dihitung berulang kali.
func climbRekursif(n int) int {
	if n <= 2 {
		return n
	}
	return climbRekursif(n-1) + climbRekursif(n-2)
}

// climbMemo menyimpan hasil yang sudah dihitung. Waktu O(n), ruang O(n).
func climbMemo(n int) int {
	memo := map[int]int{}
	var f func(int) int
	f = func(k int) int {
		if k <= 2 {
			return k
		}
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = f(k-1) + f(k-2)
		return memo[k]
	}
	return f(n)
}

// climbStairs hanya mengingat dua nilai terakhir. Waktu O(n), ruang O(1).
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

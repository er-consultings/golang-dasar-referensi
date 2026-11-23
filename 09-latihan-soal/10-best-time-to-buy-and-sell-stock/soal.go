// Package stock: prices[i] adalah harga saham di hari ke-i. Beli satu kali lalu
// jual satu kali di hari setelahnya. Kembalikan keuntungan terbesar, atau 0.
//
// Mirip LeetCode 121. Pola: satu kali jalan sambil mengingat harga terendah sejauh ini.
// Waktu O(n), ruang O(1).
package stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	termurah, untung := prices[0], 0
	for _, p := range prices[1:] {
		untung = max(untung, p-termurah) // max dan min bawaan sejak Go 1.21
		termurah = min(termurah, p)
	}
	return untung
}

// Package maxdepth: kembalikan kedalaman maksimum sebuah binary tree, yaitu jumlah
// node pada jalur terpanjang dari akar ke daun.
//
// Mirip LeetCode 104. Pola: rekursi. Kedalaman pohon = 1 + kedalaman anak yang terdalam.
// Waktu O(n), ruang O(tinggi pohon) untuk tumpukan rekursi.
package maxdepth

// TreeNode sama dengan definisi LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil { // kasus dasar: pohon kosong
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

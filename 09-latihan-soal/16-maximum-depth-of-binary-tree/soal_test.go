package maxdepth

import "testing"

func TestMaxDepth(t *testing.T) {
	//     3
	//    / \
	//   9  20
	//      / \
	//     15  7
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"contoh", root, 3},
		{"kosong", nil, 0},
		{"satu node", &TreeNode{Val: 1}, 1},
		{"miring", &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, 2},
	}
	for _, tt := range tests {
		if got := maxDepth(tt.root); got != tt.want {
			t.Errorf("%s: maxDepth = %d, ingin %d", tt.name, got, tt.want)
		}
	}
}

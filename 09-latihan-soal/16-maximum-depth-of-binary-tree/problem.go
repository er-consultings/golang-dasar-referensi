// Package maxdepth: return the maximum depth of a binary tree, that is the number of
// nodes on the longest path from the root to a leaf.
//
// Similar to LeetCode 104. Pattern: recursion. Tree depth = 1 + the depth of the deepest child.
// Time O(n), space O(tree height) for the recursion stack.
package maxdepth

// TreeNode matches the LeetCode definition.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil { // base case: empty tree
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

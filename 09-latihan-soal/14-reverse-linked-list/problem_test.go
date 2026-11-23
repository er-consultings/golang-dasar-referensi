package reverselinkedlist

import (
	"slices"
	"testing"
)

// fromSlice and toSlice make linked list tests easy to write.
func fromSlice(v []int) *ListNode {
	var head *ListNode
	for i := len(v) - 1; i >= 0; i-- {
		head = &ListNode{Val: v[i], Next: head}
	}
	return head
}

func toSlice(h *ListNode) []int {
	var out []int
	for ; h != nil; h = h.Next {
		out = append(out, h.Val)
	}
	return out
}

func TestReverseList(t *testing.T) {
	tests := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{nil, nil},
	}
	for _, tt := range tests {
		if got := toSlice(reverseList(fromSlice(tt.in))); !slices.Equal(got, tt.want) {
			t.Errorf("reverseList(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

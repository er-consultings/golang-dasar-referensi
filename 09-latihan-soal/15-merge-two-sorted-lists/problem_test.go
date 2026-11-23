package mergetwolists

import (
	"slices"
	"testing"
)

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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct{ a, b, want []int }{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, nil, nil},
		{nil, []int{0}, []int{0}},
	}
	for _, tt := range tests {
		got := toSlice(mergeTwoLists(fromSlice(tt.a), fromSlice(tt.b)))
		if !slices.Equal(got, tt.want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

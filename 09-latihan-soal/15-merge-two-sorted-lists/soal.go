// Package mergetwolists: gabungkan dua linked list terurut menjadi satu list terurut.
//
// Mirip LeetCode 21. Pola: node dummy di depan supaya tidak perlu kasus khusus untuk head.
// Waktu O(m+n), ruang O(1).
package mergetwolists

// ListNode sama dengan definisi LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	ekor := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			ekor.Next, a = a, a.Next
		} else {
			ekor.Next, b = b, b.Next
		}
		ekor = ekor.Next
	}
	if a != nil {
		ekor.Next = a
	} else {
		ekor.Next = b
	}
	return dummy.Next
}

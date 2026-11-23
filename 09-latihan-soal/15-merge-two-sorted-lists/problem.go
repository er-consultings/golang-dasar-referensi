// Package mergetwolists: merge two sorted linked lists into one sorted list.
//
// Similar to LeetCode 21. Pattern: a dummy node in front, so the head needs no special case.
// Time O(m+n), space O(1).
package mergetwolists

// ListNode matches the LeetCode definition.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

// Package reverselinkedlist: reverse a singly linked list and return its new head.
//
// Similar to LeetCode 206. Pattern: three pointers (prev, cur, next).
// Time O(n), space O(1).
package reverselinkedlist

// ListNode is exactly the definition LeetCode gives for Go problems.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // save it before the link is broken
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

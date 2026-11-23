// Package reverselinkedlist: balik urutan singly linked list dan kembalikan head barunya.
//
// Mirip LeetCode 206. Pola: tiga pointer (prev, cur, next).
// Waktu O(n), ruang O(1).
package reverselinkedlist

// ListNode sama persis dengan definisi yang diberikan LeetCode untuk soal Go.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // simpan sebelum sambungan diputus
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

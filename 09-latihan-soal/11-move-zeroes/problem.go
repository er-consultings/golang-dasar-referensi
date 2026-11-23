// Package movezeroes: move every 0 to the end without changing the order of the other
// numbers, in place.
//
// Similar to LeetCode 283. Pattern: a write pointer and a read pointer.
// Time O(n), space O(1).
package movezeroes

func moveZeroes(nums []int) {
	write := 0
	for _, n := range nums {
		if n != 0 {
			nums[write] = n
			write++
		}
	}
	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
}

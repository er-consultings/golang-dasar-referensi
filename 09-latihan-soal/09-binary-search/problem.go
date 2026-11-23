// Package binarysearch: find target in a slice sorted in ascending order. Return its index,
// or -1 if it is not there.
//
// Similar to LeetCode 704. Pattern: binary search, half of the search space is discarded at every step.
// Time O(log n), space O(1).
package binarysearch

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // a safe way to compute the middle without overflow
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

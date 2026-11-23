// Package twosum: given a slice of numbers and a target, return the indices of the two
// numbers that add up to the target. There is exactly one answer, and the same
// element may not be used twice.
//
// Similar to LeetCode 1. Pattern: a hash map stores the "numbers seen so far".
package twosum

// twoSumBrute tries every pair. Time O(n^2), space O(1).
// The first idea that comes to mind: correct, but slow for large n.
func twoSumBrute(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSum in a single pass. For every number, its complement is target - number.
// If that complement has been seen before, we have the answer.
// Time O(n), space O(n).
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums)) // number -> index
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

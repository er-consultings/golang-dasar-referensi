// Package containsduplicate: return true if any number appears
// more than once.
//
// Similar to LeetCode 217. Pattern: a map as a set.
// Time O(n), space O(n).
package containsduplicate

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

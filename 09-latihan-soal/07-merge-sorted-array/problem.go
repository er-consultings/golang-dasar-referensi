// Package mergesortedarray: nums1 holds m sorted numbers followed by n empty slots (0),
// nums2 holds n sorted numbers. Merge nums2 into nums1 so that nums1 is sorted.
//
// Similar to LeetCode 88. Pattern: fill from the back so numbers not yet moved are not overwritten.
// Time O(m+n), space O(1).
package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

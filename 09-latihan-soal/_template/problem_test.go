package template

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// Copy Examples 1, 2, 3 from the problem, then add your own edge cases:
		// an empty slice, one element, negative numbers, the largest value.
		{[]int{1, 2, 3}, -1}, // replace with Example 1 from the problem; this test fails on purpose at first
	}
	for _, tt := range tests {
		if got := solve(tt.nums); got != tt.want {
			t.Errorf("solve(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

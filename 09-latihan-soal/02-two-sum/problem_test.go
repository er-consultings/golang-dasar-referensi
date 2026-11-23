package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}
	for _, tt := range tests {
		for name, f := range map[string]func([]int, int) []int{"twoSum": twoSum, "twoSumBrute": twoSumBrute} {
			if got := f(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.target, got, tt.want)
			}
		}
	}
}

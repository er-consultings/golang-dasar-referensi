package binarysearch

import (
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	tests := []struct{ target, want int }{
		{9, 4}, {2, -1}, {-1, 0}, {12, 5}, {13, -1},
	}
	for _, tt := range tests {
		if got := search(nums, tt.target); got != tt.want {
			t.Errorf("search(%v, %d) = %d, want %d", nums, tt.target, got, tt.want)
		}
		// Compare with the standard library.
		i, ok := slices.BinarySearch(nums, tt.target)
		if ok != (tt.want >= 0) || (ok && i != tt.want) {
			t.Errorf("slices.BinarySearch differs for %d", tt.target)
		}
	}
	if search(nil, 1) != -1 {
		t.Error("an empty slice must return -1")
	}
}

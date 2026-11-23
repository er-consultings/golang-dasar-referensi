package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct{ n, want int }{
		{1, 1}, {2, 2}, {3, 3}, {5, 8}, {10, 89},
	}
	for _, tt := range tests {
		for name, f := range map[string]func(int) int{"climbStairs": climbStairs, "climbMemo": climbMemo, "climbRekursif": climbRekursif} {
			if got := f(tt.n); got != tt.want {
				t.Errorf("%s(%d) = %d, ingin %d", name, tt.n, got, tt.want)
			}
		}
	}
	if climbStairs(45) != 1836311903 {
		t.Error("climbStairs(45) salah")
	}
}

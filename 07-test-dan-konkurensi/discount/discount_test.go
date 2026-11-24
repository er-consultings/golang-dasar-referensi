package discount

import "testing"

func TestApply(t *testing.T) {
	tests := []struct {
		name    string
		total   int64
		percent int64
		want    int64
	}{
		{"no discount", 100000, 0, 100000},
		{"ten percent", 100000, 10, 90000},
		{"negative percent is ignored", 50000, -5, 50000},
		{"above 100 is capped", 50000, 150, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.total, tt.percent); got != tt.want {
				t.Errorf("Apply(%d, %d) = %d, want %d", tt.total, tt.percent, got, tt.want)
			}
		})
	}
}

// sink keeps the result alive, so the compiler cannot skip the call being measured.
var sink int64

func BenchmarkApply(b *testing.B) {
	for range b.N {
		sink = Apply(100000, 10)
	}
}

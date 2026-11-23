package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3}, {"IV", 4}, {"IX", 9}, {"LVIII", 58}, {"MCMXCIV", 1994},
	}
	for _, tt := range tests {
		if got := romanToInt(tt.s); got != tt.want {
			t.Errorf("romanToInt(%q) = %d, ingin %d", tt.s, got, tt.want)
		}
	}
}

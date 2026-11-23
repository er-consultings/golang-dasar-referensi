package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hello", "olleh"},
		{"Hannah", "hannaH"},
		{"a", "a"},
		{"", ""},
	}
	for _, tt := range tests {
		b := []byte(tt.in)
		reverseString(b)
		if string(b) != tt.want {
			t.Errorf("reverseString(%q) = %q, ingin %q", tt.in, b, tt.want)
		}
	}
}

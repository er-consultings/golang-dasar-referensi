package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
		{"Was it a car or a cat I saw?", true},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.s); got != tt.want {
			t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

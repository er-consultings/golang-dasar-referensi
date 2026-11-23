package fizzbuzz

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	got := fizzBuzz(15)
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !slices.Equal(got, want) {
		t.Errorf("fizzBuzz(15) = %v", got)
	}
	if len(fizzBuzz(0)) != 0 {
		t.Error("fizzBuzz(0) must be empty")
	}
}

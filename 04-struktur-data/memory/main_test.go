package main

import (
	"strings"
	"testing"
)

const n = 1000

// BenchmarkAppendGrow lets append grow the slice step by step.
func BenchmarkAppendGrow(b *testing.B) {
	for range b.N {
		var s []int
		for i := range n {
			s = append(s, i)
		}
	}
}

// BenchmarkAppendPrealloc reserves the capacity once with make.
func BenchmarkAppendPrealloc(b *testing.B) {
	for range b.N {
		s := make([]int, 0, n)
		for i := range n {
			s = append(s, i)
		}
	}
}

// BenchmarkConcat builds a string with +=, copying it every time.
func BenchmarkConcat(b *testing.B) {
	for range b.N {
		s := ""
		for range 100 {
			s += "go"
		}
	}
}

// BenchmarkBuilder builds the same string in one growing buffer.
func BenchmarkBuilder(b *testing.B) {
	for range b.N {
		var sb strings.Builder
		for range 100 {
			sb.WriteString("go")
		}
		_ = sb.String()
	}
}

func TestFirstThreeCopyDoesNotShare(t *testing.T) {
	big := []int{1, 2, 3, 4}
	view, copied := firstThree(big), firstThreeCopy(big)
	big[0] = 99
	if view[0] != 99 {
		t.Errorf("view[0] = %d, want 99: a view shares the array", view[0])
	}
	if copied[0] != 1 {
		t.Errorf("copied[0] = %d, want 1: a copy has its own array", copied[0])
	}
}

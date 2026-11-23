package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	s, err := Count(strings.NewReader("Go is fun.\ngo, GO!"))
	if err != nil {
		t.Fatal(err)
	}
	if s.Lines != 2 || s.Words != 5 || s.Freq["go"] != 3 {
		t.Errorf("got lines=%d words=%d go=%d, want 2, 5, 3", s.Lines, s.Words, s.Freq["go"])
	}
}

func TestRunFromStdin(t *testing.T) {
	var out bytes.Buffer
	err := run([]string{"-top", "1"}, strings.NewReader("a b b"), &out)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "b               2") {
		t.Errorf("output:\n%s", out.String())
	}
}

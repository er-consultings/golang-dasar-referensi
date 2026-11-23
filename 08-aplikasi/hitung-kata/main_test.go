package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHitung(t *testing.T) {
	s, err := Hitung(strings.NewReader("Go itu asyik.\ngo, GO!"))
	if err != nil {
		t.Fatal(err)
	}
	if s.Baris != 2 || s.Kata != 5 || s.Frek["go"] != 3 {
		t.Errorf("got baris=%d kata=%d go=%d, ingin 2, 5, 3", s.Baris, s.Kata, s.Frek["go"])
	}
}

func TestRunDariStdin(t *testing.T) {
	var out bytes.Buffer
	err := run([]string{"-top", "1"}, strings.NewReader("a b b"), &out)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "b               2") {
		t.Errorf("output:\n%s", out.String())
	}
}

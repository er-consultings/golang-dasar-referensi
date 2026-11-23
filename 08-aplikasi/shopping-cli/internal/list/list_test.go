package list

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestAddAndTotal(t *testing.T) {
	var l List // the zero value is used directly
	if err := l.Add("coffee", 2, 18000); err != nil {
		t.Fatal(err)
	}
	if err := l.Add("coffee", 1, 18000); err != nil {
		t.Fatal(err)
	}
	if len(l.Items) != 1 {
		t.Fatalf("items = %d lines, want 1", len(l.Items))
	}
	if got, want := l.Total(), int64(54000); got != want {
		t.Errorf("Total() = %d, want %d", got, want)
	}
}

func TestAddInvalid(t *testing.T) {
	var l List
	if err := l.Add("bread", 0, 15000); !errors.Is(err, ErrInvalidQty) {
		t.Errorf("err = %v, want ErrInvalidQty", err)
	}
}

func TestRemove(t *testing.T) {
	var l List
	_ = l.Add("coffee", 1, 18000)
	_ = l.Add("bread", 1, 15000)
	if err := l.Remove("coffee"); err != nil {
		t.Fatal(err)
	}
	if len(l.Items) != 1 || l.Items[0].Name != "bread" {
		t.Errorf("remaining = %+v", l.Items)
	}
	if err := l.Remove("tea"); !errors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestSaveThenLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "shopping.json") // temporary folder, removed automatically
	var l List
	_ = l.Add("coffee", 2, 18000)
	if err := l.Save(path); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.Total() != 36000 {
		t.Errorf("Total after loading = %d", got.Total())
	}
}

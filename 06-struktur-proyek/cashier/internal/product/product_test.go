package product

import (
	"errors"
	"testing"
)

func TestFind(t *testing.T) {
	c := NewCatalog([]Product{{ID: "coffee", Name: "Coffee", Price: 18000}})
	p, err := c.Find("coffee")
	if err != nil || p.Price != 18000 {
		t.Fatalf("Find(coffee) = %+v, %v", p, err)
	}
	if _, err := c.Find("tea"); !errors.Is(err, ErrNotFound) {
		t.Errorf("Find(tea) err = %v, want ErrNotFound", err)
	}
}

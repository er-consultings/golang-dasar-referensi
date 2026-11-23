package order

import (
	"errors"
	"testing"

	"golang-dasar-referensi/06-struktur-proyek/cashier/internal/product"
)

// fakeCatalog satisfies the Catalog interface without the real product package.
type fakeCatalog map[string]product.Product

func (c fakeCatalog) Find(id string) (product.Product, error) {
	p, ok := c[id]
	if !ok {
		return product.Product{}, product.ErrNotFound
	}
	return p, nil
}

func TestCreate(t *testing.T) {
	svc := NewService(fakeCatalog{"coffee": {ID: "coffee", Name: "Coffee", Price: 18000}})

	tests := []struct {
		name      string
		lines     []Line
		wantTotal int64
		wantErr   error
	}{
		{"one product", []Line{{"coffee", 2}}, 36000, nil},
		{"empty", nil, 0, ErrEmptyOrder},
		{"not in catalog", []Line{{"tea", 1}}, 0, product.ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o, err := svc.Create(tt.lines)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("err = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := o.Total(); got != tt.wantTotal {
				t.Errorf("Total() = %d, want %d", got, tt.wantTotal)
			}
		})
	}
}

// Package order creates orders and calculates their totals.
package order

import (
	"errors"
	"fmt"

	"golang-dasar-referensi/06-struktur-proyek/cashier/internal/product"
)

// ErrEmptyOrder is returned when an order contains no items.
var ErrEmptyOrder = errors.New("empty order")

// Catalog is the only thing this package needs from the product catalog.
// The interface is defined on the consumer side, so tests can pass a fake.
type Catalog interface {
	Find(id string) (product.Product, error)
}

// Line is a request for one product and its quantity.
type Line struct {
	ProductID string
	Qty       int
}

// Item is one order line that already has a name and a price.
type Item struct {
	Name  string
	Price int64
	Qty   int
}

// Subtotal returns the price times the quantity.
func (it Item) Subtotal() int64 { return it.Price * int64(it.Qty) }

// Order holds the priced items.
type Order struct {
	Items []Item
}

// Total adds up every subtotal.
func (o Order) Total() int64 {
	var t int64
	for _, it := range o.Items {
		t += it.Subtotal()
	}
	return t
}

// Service creates orders using a Catalog.
type Service struct {
	catalog Catalog
}

// NewService creates a Service.
func NewService(c Catalog) *Service { return &Service{catalog: c} }

// Create checks every line and fills in the name and price from the catalog.
func (s *Service) Create(lines []Line) (Order, error) {
	if len(lines) == 0 {
		return Order{}, ErrEmptyOrder
	}
	out := Order{Items: make([]Item, 0, len(lines))}
	for _, l := range lines {
		if l.Qty <= 0 {
			return Order{}, fmt.Errorf("quantity of %s must be greater than zero", l.ProductID)
		}
		p, err := s.catalog.Find(l.ProductID)
		if err != nil {
			return Order{}, fmt.Errorf("product %s: %w", l.ProductID, err)
		}
		out.Items = append(out.Items, Item{Name: p.Name, Price: p.Price, Qty: l.Qty})
	}
	return out, nil
}

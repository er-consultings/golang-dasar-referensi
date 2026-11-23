// Package product stores the catalog of products for sale.
package product

import "errors"

// ErrNotFound is returned when a product ID is not in the catalog.
var ErrNotFound = errors.New("product not found")

// Product is an item for sale. The price is in rupiah.
type Product struct {
	ID    string
	Name  string
	Price int64
}

// Catalog keeps products in memory.
type Catalog struct {
	byID map[string]Product
}

// NewCatalog creates a catalog from a list of products.
func NewCatalog(products []Product) *Catalog {
	c := &Catalog{byID: make(map[string]Product, len(products))}
	for _, p := range products {
		c.byID[p.ID] = p
	}
	return c
}

// Find returns the product with the given ID.
func (c *Catalog) Find(id string) (Product, error) {
	p, ok := c.byID[id]
	if !ok {
		return Product{}, ErrNotFound
	}
	return p, nil
}

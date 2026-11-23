// Package list stores shopping items and calculates their total.
package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ErrInvalidQty is returned when the quantity or price is not positive.
var ErrInvalidQty = errors.New("quantity and price must be greater than zero")

// ErrNotFound is returned when removing an item that is not on the list.
var ErrNotFound = errors.New("item is not on the list")

// Item is one line on the shopping list. The price is in rupiah.
type Item struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price int64  `json:"price"`
}

// Subtotal returns the price times the quantity.
func (it Item) Subtotal() int64 { return it.Price * int64(it.Qty) }

// List is a collection of items. Its zero value is ready to use.
type List struct {
	Items []Item `json:"items"`
}

// Add inserts a new item, or increases the quantity when the name already exists.
func (l *List) Add(name string, qty int, price int64) error {
	if qty <= 0 || price <= 0 {
		return ErrInvalidQty
	}
	for i := range l.Items {
		if l.Items[i].Name == name {
			l.Items[i].Qty += qty
			l.Items[i].Price = price
			return nil
		}
	}
	l.Items = append(l.Items, Item{Name: name, Qty: qty, Price: price})
	return nil
}

// Remove deletes the item with the given name.
func (l *List) Remove(name string) error {
	for i := range l.Items {
		if l.Items[i].Name == name {
			l.Items = append(l.Items[:i], l.Items[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// Total adds up the subtotals of every item.
func (l *List) Total() int64 {
	var total int64
	for _, it := range l.Items {
		total += it.Subtotal()
	}
	return total
}

// Load reads a list from a JSON file. A file that does not exist yet gives an empty list.
func Load(path string) (*List, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return &List{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var l List
	if err := json.Unmarshal(data, &l); err != nil {
		return nil, fmt.Errorf("invalid format in %s: %w", path, err)
	}
	return &l, nil
}

// Save writes the list to a JSON file.
func (l *List) Save(path string) error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

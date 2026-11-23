// Errors as values: sentinels, wrapping with %w, errors.Is, errors.As.
//
//	go run ./03-fungsi-error/errors
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error that callers can recognize.
var ErrNotFound = errors.New("data not found")

// ValidationError carries the name of the invalid field.
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string { return "invalid field: " + e.Field }

type Product struct {
	ID   int
	Name string
}

var warehouse = map[int]Product{1: {ID: 1, Name: "Coffee"}}

func find(id int) (Product, error) {
	p, ok := warehouse[id]
	if !ok {
		return Product{}, ErrNotFound
	}
	return p, nil
}

func loadProduct(id int) (Product, error) {
	if id <= 0 {
		return Product{}, &ValidationError{Field: "id"}
	}
	p, err := find(id)
	if err != nil {
		return Product{}, fmt.Errorf("loading product %d: %w", id, err)
	}
	return p, nil
}

func main() {
	for _, id := range []int{1, 7, -1} {
		p, err := loadProduct(id)
		var ve *ValidationError
		switch {
		case err == nil:
			fmt.Println("found:", p.Name)
		case errors.Is(err, ErrNotFound):
			fmt.Println("404:", err)
		case errors.As(err, &ve):
			fmt.Println("400, check field", ve.Field)
		default:
			fmt.Println("other error:", err)
		}
	}
}

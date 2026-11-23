// Pointers: & and *, changing the caller's value, nil pointers.
//
//	go run ./04-struktur-data/pointer
package main

import "fmt"

type Product struct {
	Name  string
	Price int64
}

func applyDiscount(p *Product, percent int64) {
	p.Price -= p.Price * percent / 100
}

func newProduct() *Product {
	p := Product{Name: "New"}
	return &p // safe, the compiler moves it to the heap
}

func main() {
	x := 10
	p := &x
	*p = 20
	fmt.Println("x:", x)

	coffee := Product{Name: "Coffee", Price: 50000}
	applyDiscount(&coffee, 10)
	fmt.Println("price after discount:", coffee.Price)

	fmt.Println(newProduct().Name)

	var empty *Product
	if empty == nil {
		fmt.Println("nil pointer, do not access its fields")
	}
}

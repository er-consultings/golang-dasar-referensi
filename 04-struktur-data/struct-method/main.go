// Structs, constructors, methods with value and pointer receivers, JSON tags.
//
//	go run ./04-struktur-data/struct-method
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int    `json:"stock"`
	cost  int64  // lowercase: not included in JSON
}

func NewProduct(name string, price int64) (*Product, error) {
	if price < 0 {
		return nil, errors.New("price must not be negative")
	}
	return &Product{Name: name, Price: price}, nil
}

func (p *Product) Total(qty int) int64 { return p.Price * int64(qty) }

func (p *Product) ReduceStock(n int) error {
	if n > p.Stock {
		return fmt.Errorf("only %d %s left in stock", p.Stock, p.Name)
	}
	p.Stock -= n
	return nil
}

func main() {
	p, err := NewProduct("Arabica Coffee", 85000)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Stock = 3
	p.cost = 60000
	fmt.Printf("%+v\n", *p)
	fmt.Println("total for 2:", p.Total(2))

	if err := p.ReduceStock(5); err != nil {
		fmt.Println("failed:", err)
	}
	_ = p.ReduceStock(2)
	fmt.Println("stock left:", p.Stock)

	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	var fromJSON Product
	if err := json.Unmarshal([]byte(`{"id":2,"name":"Tea","price":12000}`), &fromJSON); err == nil {
		fmt.Printf("%+v\n", fromJSON)
	}
}

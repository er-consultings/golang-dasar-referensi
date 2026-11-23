// Composition with embedding, promoted methods, a named mutex.
//
//	go run ./05-interface/composition
package main

import (
	"fmt"
	"sync"
)

type Address struct {
	City       string
	PostalCode string
}

func (a Address) Full() string { return a.City + " " + a.PostalCode }

type Customer struct {
	Name string
	Address
}

type Counter struct {
	mu sync.Mutex // named, so Lock and Unlock are not visible from outside
	n  map[string]int
}

func NewCounter() *Counter { return &Counter{n: map[string]int{}} }

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n[key]++
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n[key]
}

func main() {
	cust := Customer{Name: "Rina", Address: Address{City: "Malang", PostalCode: "65145"}}
	fmt.Println(cust.City, "|", cust.Full())

	counter := NewCounter()
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc("click")
		}()
	}
	wg.Wait()
	fmt.Println("clicks:", counter.Value("click"))
}

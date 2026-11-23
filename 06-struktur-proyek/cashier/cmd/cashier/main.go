// Command cashier creates one sample order and prints its receipt.
package main

import (
	"fmt"
	"os"

	"golang-dasar-referensi/06-struktur-proyek/cashier/internal/order"
	"golang-dasar-referensi/06-struktur-proyek/cashier/internal/product"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// run wires the dependencies together. The logic lives in the internal packages.
func run() error {
	catalog := product.NewCatalog([]product.Product{
		{ID: "coffee", Name: "Milk Coffee", Price: 18000},
		{ID: "toast", Name: "Toast", Price: 15000},
	})
	svc := order.NewService(catalog)

	o, err := svc.Create([]order.Line{{ProductID: "coffee", Qty: 2}, {ProductID: "toast", Qty: 1}})
	if err != nil {
		return fmt.Errorf("creating order: %w", err)
	}
	for _, it := range o.Items {
		fmt.Printf("%-12s %d x %d = %d\n", it.Name, it.Qty, it.Price, it.Subtotal())
	}
	fmt.Println("Total:", o.Total())
	return nil
}

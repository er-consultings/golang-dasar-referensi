// Command kasir membuat satu pesanan contoh dan mencetak struknya.
package main

import (
	"fmt"
	"os"

	"golang-dasar-referensi/06-struktur-proyek/kasir/internal/pesanan"
	"golang-dasar-referensi/06-struktur-proyek/kasir/internal/produk"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// run merangkai dependensi. Logika ada di paket internal.
func run() error {
	katalog := produk.NewKatalog([]produk.Produk{
		{ID: "kopi", Nama: "Kopi Susu", Harga: 18000},
		{ID: "roti", Nama: "Roti Bakar", Harga: 15000},
	})
	layanan := pesanan.NewLayanan(katalog)

	p, err := layanan.Buat([]pesanan.Baris{{ProdukID: "kopi", Jumlah: 2}, {ProdukID: "roti", Jumlah: 1}})
	if err != nil {
		return fmt.Errorf("membuat pesanan: %w", err)
	}
	for _, b := range p.Baris {
		fmt.Printf("%-10s %d x %d = %d\n", b.Nama, b.Jumlah, b.Harga, b.Subtotal())
	}
	fmt.Println("Total:", p.Total())
	return nil
}

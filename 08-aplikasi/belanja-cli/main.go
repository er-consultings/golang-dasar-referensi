// Command belanja mencatat daftar belanja di berkas belanja.json.
//
//	go run ./08-aplikasi/belanja-cli tambah kopi 2 18000
//	go run ./08-aplikasi/belanja-cli daftar
//	go run ./08-aplikasi/belanja-cli hapus kopi
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"golang-dasar-referensi/08-aplikasi/belanja-cli/internal/daftar"
)

const berkas = "belanja.json"

const pemakaian = "pakai: belanja tambah <nama> <jumlah> <harga> | belanja hapus <nama> | belanja daftar"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New(pemakaian)
	}
	d, err := daftar.Muat(berkas)
	if err != nil {
		return err
	}

	switch args[0] {
	case "tambah":
		if len(args) != 4 {
			return errors.New(pemakaian)
		}
		jumlah, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("jumlah %q bukan angka", args[2])
		}
		harga, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return fmt.Errorf("harga %q bukan angka", args[3])
		}
		if err := d.Tambah(args[1], jumlah, harga); err != nil {
			return err
		}
		return d.Simpan(berkas)
	case "hapus":
		if len(args) != 2 {
			return errors.New(pemakaian)
		}
		if err := d.Hapus(args[1]); err != nil {
			return err
		}
		return d.Simpan(berkas)
	case "daftar":
		for _, b := range d.Barang {
			fmt.Printf("%-6s %d x %d = %d\n", b.Nama, b.Jumlah, b.Harga, b.Subtotal())
		}
		fmt.Println("Total:", d.Total())
		return nil
	default:
		return fmt.Errorf("perintah tidak dikenal: %s", args[0])
	}
}

// defer (urutan terbalik), panic, dan recover.
//
//	go run ./03-fungsi-error/defer-panic
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tulisSementara() error {
	path := filepath.Join(os.TempDir(), "contoh-defer.txt")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer os.Remove(path) // dijalankan terakhir
	defer f.Close()       // dijalankan lebih dulu

	_, err = fmt.Fprintln(f, "halo dari defer")
	return err
}

// aman mengubah panic menjadi error biasa.
func aman(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("terjadi panic: %v", r)
		}
	}()
	f()
	return nil
}

func main() {
	for i := range 3 {
		defer fmt.Println("defer", i) // tercetak 2, 1, 0 saat main selesai
	}

	if err := tulisSementara(); err != nil {
		fmt.Println("gagal:", err)
	}

	err := aman(func() {
		var angka []int
		_ = angka[5] // indeks di luar batas, panic
	})
	fmt.Println(err)
}

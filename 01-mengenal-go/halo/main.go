// Program pertama. Jalankan dari akar repositori:
//
//	go run ./01-mengenal-go/halo
//	go run ./01-mengenal-go/halo Sari
package main

import (
	"fmt"
	"os"
)

func main() {
	nama := "Budi"
	if len(os.Args) > 1 {
		nama = os.Args[1] // argumen pertama setelah nama program
	}
	fmt.Println("Halo,", nama)
	fmt.Printf("Nama Anda terdiri dari %d huruf\n", len(nama))
	fmt.Printf("Tipe variabel nama: %T\n", nama)
}

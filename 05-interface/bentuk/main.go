// Interface implisit, interface kecil, io.Reader, type switch, jebakan nil.
//
//	go run ./05-interface/bentuk
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type Bentuk interface {
	Luas() float64
}

type Persegi struct{ Sisi float64 }
type Lingkaran struct{ R float64 }

func (p Persegi) Luas() float64   { return p.Sisi * p.Sisi }
func (l Lingkaran) Luas() float64 { return math.Pi * l.R * l.R }

// String memenuhi fmt.Stringer, sehingga Println memakai teks ini.
func (p Persegi) String() string { return fmt.Sprintf("Persegi %.0f", p.Sisi) }

func totalLuas(daftar []Bentuk) float64 {
	total := 0.0
	for _, b := range daftar {
		total += b.Luas()
	}
	return total
}

// hitungBaris menerima io.Reader, jadi bisa dipakai untuk berkas, stdin, atau string.
func hitungBaris(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	n := 0
	for sc.Scan() {
		n++
	}
	return n, sc.Err()
}

func jelaskan(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("angka %d", x)
	case string:
		return fmt.Sprintf("teks %q", x)
	case Bentuk:
		return fmt.Sprintf("bentuk dengan luas %.2f", x.Luas())
	default:
		return fmt.Sprintf("tipe lain %T", x)
	}
}

type ErrKhusus struct{}

func (*ErrKhusus) Error() string { return "khusus" }

func salah() error {
	var e *ErrKhusus // nil
	return e         // interface error ini TIDAK nil
}

func benar() error { return nil }

func main() {
	fmt.Printf("%.2f\n", totalLuas([]Bentuk{Persegi{2}, Lingkaran{1}}))
	fmt.Println(Persegi{3})

	n, _ := hitungBaris(strings.NewReader("a\nb\nc"))
	fmt.Println("baris:", n)

	for _, v := range []any{7, "go", Lingkaran{2}, 3.5} {
		fmt.Println(jelaskan(v))
	}

	fmt.Println("salah() != nil:", salah() != nil)
	fmt.Println("benar() != nil:", benar() != nil)
}

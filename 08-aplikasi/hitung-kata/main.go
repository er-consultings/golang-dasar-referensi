// Command hitung-kata menghitung baris, kata, dan kata yang paling sering muncul
// dari sebuah berkas atau dari input standar.
//
//	go run ./08-aplikasi/hitung-kata README.md
//	echo "satu dua dua tiga tiga tiga" | go run ./08-aplikasi/hitung-kata -top 2
package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"unicode"
)

// Statistik adalah hasil hitungan sebuah teks.
type Statistik struct {
	Baris int
	Kata  int
	Frek  map[string]int
}

// Hitung membaca r sampai habis. Huruf besar dan kecil dianggap sama,
// tanda baca di awal dan akhir kata dibuang.
func Hitung(r io.Reader) (Statistik, error) {
	s := Statistik{Frek: map[string]int{}}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		s.Baris++
		for _, w := range strings.Fields(sc.Text()) {
			w = strings.ToLower(strings.TrimFunc(w, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsDigit(r)
			}))
			if w == "" {
				continue
			}
			s.Kata++
			s.Frek[w]++
		}
	}
	return s, sc.Err()
}

type pasangan struct {
	kata   string
	jumlah int
}

// Teratas mengembalikan n kata paling sering, seri diurutkan abjad.
func (s Statistik) Teratas(n int) []pasangan {
	out := make([]pasangan, 0, len(s.Frek))
	for k, v := range s.Frek {
		out = append(out, pasangan{k, v})
	}
	slices.SortFunc(out, func(a, b pasangan) int {
		if c := cmp.Compare(b.jumlah, a.jumlah); c != 0 {
			return c
		}
		return cmp.Compare(a.kata, b.kata)
	})
	return out[:min(n, len(out))]
}

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	fs := flag.NewFlagSet("hitung-kata", flag.ContinueOnError)
	top := fs.Int("top", 5, "jumlah kata teratas yang ditampilkan")
	if err := fs.Parse(args); err != nil {
		return err
	}

	in := stdin
	if fs.NArg() > 0 {
		f, err := os.Open(fs.Arg(0))
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	s, err := Hitung(in)
	if err != nil {
		return fmt.Errorf("membaca input: %w", err)
	}
	fmt.Fprintf(stdout, "baris: %d, kata: %d, kata unik: %d\n", s.Baris, s.Kata, len(s.Frek))
	for _, p := range s.Teratas(*top) {
		fmt.Fprintf(stdout, "%-15s %d\n", p.kata, p.jumlah)
	}
	return nil
}

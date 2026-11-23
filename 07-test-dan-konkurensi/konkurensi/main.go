// Goroutine, WaitGroup, channel untuk mengumpulkan hasil, dan Mutex.
//
//	go run ./07-test-dan-konkurensi/konkurensi
//	go run -race ./07-test-dan-konkurensi/konkurensi
package main

import (
	"fmt"
	"sync"
	"time"
)

// ukur berpura-pura melakukan pekerjaan lambat, misalnya mengunduh halaman.
func ukur(url string) int {
	time.Sleep(50 * time.Millisecond)
	return len(url) * 100
}

func ukurSemua(urls []string) map[string]int {
	type hasil struct {
		url    string
		ukuran int
	}
	ch := make(chan hasil, len(urls))
	for _, u := range urls {
		go func() {
			ch <- hasil{u, ukur(u)}
		}()
	}
	out := make(map[string]int, len(urls))
	for range urls {
		h := <-ch // hanya goroutine ini yang menulis ke map
		out[h.url] = h.ukuran
	}
	return out
}

type Penghitung struct {
	mu sync.Mutex
	n  int
}

func (p *Penghitung) Tambah() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.n++
}

func main() {
	mulai := time.Now()
	fmt.Println(ukurSemua([]string{"a.com", "bb.com", "ccc.com"}))
	fmt.Println("selesai dalam", time.Since(mulai).Round(10*time.Millisecond), "bukan 150ms, karena berjalan bersamaan")

	var wg sync.WaitGroup
	var p Penghitung
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.Tambah()
		}()
	}
	wg.Wait()
	fmt.Println("penghitung:", p.n)

	// Pengirim yang menutup channel, penerima membaca sampai habis dengan range.
	angka := make(chan int)
	go func() {
		defer close(angka)
		for i := range 5 {
			angka <- i * i
		}
	}()
	for n := range angka {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

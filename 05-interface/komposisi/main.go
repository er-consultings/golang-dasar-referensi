// Komposisi dengan embedding, method yang dipromosikan, mutex bernama.
//
//	go run ./05-interface/komposisi
package main

import (
	"fmt"
	"sync"
)

type Alamat struct {
	Kota    string
	KodePos string
}

func (a Alamat) Lengkap() string { return a.Kota + " " + a.KodePos }

type Pelanggan struct {
	Nama string
	Alamat
}

type Penghitung struct {
	mu sync.Mutex // bernama, Lock dan Unlock tidak ikut terlihat dari luar
	n  map[string]int
}

func NewPenghitung() *Penghitung { return &Penghitung{n: map[string]int{}} }

func (p *Penghitung) Tambah(key string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.n[key]++
}

func (p *Penghitung) Nilai(key string) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.n[key]
}

func main() {
	c := Pelanggan{Nama: "Rina", Alamat: Alamat{Kota: "Malang", KodePos: "65145"}}
	fmt.Println(c.Kota, "|", c.Lengkap())

	h := NewPenghitung()
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.Tambah("klik")
		}()
	}
	wg.Wait()
	fmt.Println("klik:", h.Nilai("klik"))
}

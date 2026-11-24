// Data races and how sync.Mutex, sync.RWMutex, and sync/atomic prevent them.
//
//	go run ./07-test-dan-konkurensi/mutex
//	go run -race ./07-test-dan-konkurensi/mutex -unsafe   # watch the race detector catch it
//	go test -race ./07-test-dan-konkurensi/mutex
package main

import (
	"flag"
	"fmt"
	"sync"
)

const workers = 1000

func main() {
	unsafe := flag.Bool("unsafe", false, "also run the counter without a lock")
	flag.Parse()

	if *unsafe {
		var c UnsafeCounter
		var wg sync.WaitGroup
		for range workers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}
		wg.Wait()
		fmt.Printf("unsafe counter: %d (expected %d)\n", c.n, workers)
	}

	var acc Account
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(1000)
		}()
	}
	wg.Wait()
	fmt.Println("balance after 1000 deposits:", acc.Balance())

	// 1500 withdrawals of 1000 against a balance of 1,000,000: exactly 1000 succeed.
	var ok, failed Visits
	for range 1500 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := acc.Withdraw(1000); err != nil {
				failed.Inc()
				return
			}
			ok.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("withdrawals: %d ok, %d refused, balance %d\n", ok.Count(), failed.Count(), acc.Balance())

	prices := NewPriceList()
	prices.Set("coffee", 18000)
	v, found := prices.Get("coffee")
	fmt.Println("coffee:", v, found)
}

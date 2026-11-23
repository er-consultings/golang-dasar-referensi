// Goroutines, WaitGroup, channels for collecting results, and Mutex.
//
//	go run ./07-test-dan-konkurensi/concurrency
//	go run -race ./07-test-dan-konkurensi/concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

// measure pretends to do slow work, such as downloading a page.
func measure(url string) int {
	time.Sleep(50 * time.Millisecond)
	return len(url) * 100
}

func measureAll(urls []string) map[string]int {
	type result struct {
		url  string
		size int
	}
	ch := make(chan result, len(urls))
	for _, u := range urls {
		go func() {
			ch <- result{u, measure(u)}
		}()
	}
	out := make(map[string]int, len(urls))
	for range urls {
		r := <-ch // only this goroutine writes to the map
		out[r.url] = r.size
	}
	return out
}

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func main() {
	start := time.Now()
	fmt.Println(measureAll([]string{"a.com", "bb.com", "ccc.com"}))
	fmt.Println("done in", time.Since(start).Round(10*time.Millisecond), "not 150ms, because they ran concurrently")

	var wg sync.WaitGroup
	var c Counter
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", c.n)

	// The sender closes the channel; the receiver reads until it is drained with range.
	squares := make(chan int)
	go func() {
		defer close(squares)
		for i := range 5 {
			squares <- i * i
		}
	}()
	for n := range squares {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

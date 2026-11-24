// context: deadlines, cancellation, and stopping goroutines.
//
//	go run ./07-test-dan-konkurensi/context
//	go run ./07-test-dan-konkurensi/context -timeout 2s
//	go run ./07-test-dan-konkurensi/context -timeout 1m   # then press Ctrl+C
//	go test -race ./07-test-dan-konkurensi/context
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	timeout := flag.Duration("timeout", 350*time.Millisecond, "how long the search may take")
	flag.Parse()

	// ctx is cancelled when the user presses Ctrl+C.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// A deadline on top of that: whichever comes first ends the search.
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	items := []string{"go basics", "python basics", "go testing", "go context", "rust basics", "go modules"}
	found, err := Search(ctx, items, "go", 100*time.Millisecond)
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		fmt.Println("too slow, partial results:", found)
	case errors.Is(err, context.Canceled):
		fmt.Println("cancelled with Ctrl+C, partial results:", found)
	case err != nil:
		return err
	default:
		fmt.Println("all results:", found)
	}

	// Take five numbers, then cancel so the generator goroutine can finish.
	genCtx, cancelGen := context.WithCancel(context.Background())
	numbers := Generate(genCtx)
	for range 5 {
		fmt.Print(<-numbers, " ")
	}
	fmt.Println()
	cancelGen()
	for range numbers { // drain until the generator closes the channel
	}
	fmt.Println("generator stopped")
	return nil
}

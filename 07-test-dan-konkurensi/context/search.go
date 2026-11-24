package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Search pretends to look through slow storage, one item per step. It checks ctx
// before every step and stops as soon as the caller no longer wants the answer.
func Search(ctx context.Context, items []string, word string, step time.Duration) ([]string, error) {
	var found []string
	for _, it := range items {
		select {
		case <-ctx.Done():
			return found, fmt.Errorf("search %q stopped after %d results: %w", word, len(found), ctx.Err())
		case <-time.After(step): // the slow work for one item
		}
		if strings.Contains(it, word) {
			found = append(found, it)
		}
	}
	return found, nil
}

// Generate sends 1, 2, 3, ... until ctx is cancelled, then closes the channel.
// Without ctx this goroutine would block forever once nobody reads: a goroutine leak.
func Generate(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

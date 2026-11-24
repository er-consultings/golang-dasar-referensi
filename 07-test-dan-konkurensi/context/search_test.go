package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSearchFinishesInTime(t *testing.T) {
	got, err := Search(context.Background(), []string{"go a", "b", "go c"}, "go", time.Millisecond)
	if err != nil || len(got) != 2 {
		t.Fatalf("Search = %v, %v; want 2 results and no error", got, err)
	}
}

func TestSearchStopsAtDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	start := time.Now()
	items := make([]string, 100) // 100 steps of 10ms would take a second
	_, err := Search(ctx, items, "go", 10*time.Millisecond)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want context.DeadlineExceeded", err)
	}
	if took := time.Since(start); took > 500*time.Millisecond {
		t.Errorf("Search took %v after the deadline passed", took)
	}
}

func TestSearchAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Search(ctx, []string{"go"}, "go", time.Millisecond); !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestGenerateStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := Generate(ctx)
	if first := <-numbers; first != 1 {
		t.Fatalf("first number = %d, want 1", first)
	}
	cancel()
	done := make(chan struct{})
	go func() {
		for range numbers {
		}
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("the generator did not close its channel after cancel")
	}
}

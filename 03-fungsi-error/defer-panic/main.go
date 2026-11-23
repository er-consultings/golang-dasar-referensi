// defer (reverse order), panic, and recover.
//
//	go run ./03-fungsi-error/defer-panic
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeTemp() error {
	path := filepath.Join(os.TempDir(), "defer-example.txt")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer os.Remove(path) // runs last
	defer f.Close()       // runs first

	_, err = fmt.Fprintln(f, "hello from defer")
	return err
}

// safely turns a panic into an ordinary error.
func safely(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	f()
	return nil
}

func main() {
	for i := range 3 {
		defer fmt.Println("defer", i) // prints 2, 1, 0 when main returns
	}

	if err := writeTemp(); err != nil {
		fmt.Println("failed:", err)
	}

	err := safely(func() {
		var numbers []int
		_ = numbers[5] // index out of range, panics
	})
	fmt.Println(err)
}

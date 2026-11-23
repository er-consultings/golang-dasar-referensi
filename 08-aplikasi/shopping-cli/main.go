// Command shopping keeps a shopping list in the file shopping.json.
//
//	go run ./08-aplikasi/shopping-cli add coffee 2 18000
//	go run ./08-aplikasi/shopping-cli list
//	go run ./08-aplikasi/shopping-cli remove coffee
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"golang-dasar-referensi/08-aplikasi/shopping-cli/internal/list"
)

const dataFile = "shopping.json"

const usage = "usage: shopping add <name> <qty> <price> | shopping remove <name> | shopping list"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New(usage)
	}
	l, err := list.Load(dataFile)
	if err != nil {
		return err
	}

	switch args[0] {
	case "add":
		if len(args) != 4 {
			return errors.New(usage)
		}
		qty, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("quantity %q is not a number", args[2])
		}
		price, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return fmt.Errorf("price %q is not a number", args[3])
		}
		if err := l.Add(args[1], qty, price); err != nil {
			return err
		}
		return l.Save(dataFile)
	case "remove":
		if len(args) != 2 {
			return errors.New(usage)
		}
		if err := l.Remove(args[1]); err != nil {
			return err
		}
		return l.Save(dataFile)
	case "list":
		for _, it := range l.Items {
			fmt.Printf("%-6s %d x %d = %d\n", it.Name, it.Qty, it.Price, it.Subtotal())
		}
		fmt.Println("Total:", l.Total())
		return nil
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

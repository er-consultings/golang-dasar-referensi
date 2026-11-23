// First program. Run it from the repository root:
//
//	go run ./01-mengenal-go/hello
//	go run ./01-mengenal-go/hello Sari
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Budi"
	if len(os.Args) > 1 {
		name = os.Args[1] // first argument after the program name
	}
	fmt.Println("Hello,", name)
	fmt.Printf("Your name has %d letters\n", len(name))
	fmt.Printf("Type of name: %T\n", name)
}

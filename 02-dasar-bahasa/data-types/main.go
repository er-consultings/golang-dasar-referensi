// Basic data types, string, byte, rune, and conversions.
//
//	go run ./02-dasar-bahasa/data-types
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	a, b := 7, 2
	fmt.Println("7 / 2 =", a/b)                           // 3
	fmt.Println("7 % 2 =", a%b)                           // 1
	fmt.Println("float:", float64(a)/float64(b))          // 3.5
	fmt.Println("0.1 + 0.2 =", 0.1+0.2 == 0.3, "(float)") // be careful with float

	// Money is stored as a whole number of rupiah.
	var price int64 = 18000
	fmt.Println("total for 3 coffees:", price*3)

	k := "Coffee ☕"
	fmt.Println("bytes:", len(k), "characters:", utf8.RuneCountInString(k))
	for i, r := range k {
		fmt.Printf("  index %d: %q\n", i, r)
	}

	// Converting text to a number can always fail, so check the error.
	for _, input := range []string{"42", "42a"} {
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("failed:", err)
			continue
		}
		fmt.Println("number:", n)
	}
	fmt.Println("Itoa:", strconv.Itoa(65), "not", string(rune(65)))

	raw := `first line
second line \n is not processed`
	fmt.Println(raw)
}

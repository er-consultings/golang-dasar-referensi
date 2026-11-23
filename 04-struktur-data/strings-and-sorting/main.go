// Tools that show up often in practice problems: strings, strings.Builder,
// byte and rune, unicode, and sorting with slices and sort.
//
//	go run ./04-struktur-data/strings-and-sorting
package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	s := "  Hello, Go World  "
	fmt.Printf("%q\n", strings.TrimSpace(s))
	fmt.Println(strings.ToLower(s), strings.ToUpper(s))
	fmt.Println(strings.Contains(s, "Go"), strings.Index(s, "World"))
	fmt.Println(strings.Split("a,b,c", ","), strings.Fields(s))
	fmt.Println(strings.Join([]string{"x", "y"}, "-"), strings.Repeat("ab", 3))

	// Building a long string: use strings.Builder, not repeated +.
	var b strings.Builder
	for i := range 5 {
		fmt.Fprintf(&b, "%d;", i)
	}
	fmt.Println(b.String())

	// Strings are immutable. Convert to []byte or []rune first.
	word := []rune("stressed")
	slices.Reverse(word)
	fmt.Println(string(word))

	// Bytes as numbers: counting the letters a to z.
	var counts [26]int
	for _, c := range "banana" {
		counts[c-'a']++
	}
	fmt.Println("a:", counts['a'-'a'], "n:", counts['n'-'a'])

	fmt.Println(unicode.IsLetter('x'), unicode.IsDigit('7'), string(unicode.ToLower('Q')))

	numbers := []int{5, 2, 9, 1}
	slices.Sort(numbers)
	fmt.Println(numbers)

	students := []Student{{"Rina", 88}, {"Budi", 92}, {"Sari", 88}}
	// Highest score first, then name from A to Z.
	slices.SortFunc(students, func(a, b Student) int {
		if c := cmp.Compare(b.Score, a.Score); c != 0 {
			return c
		}
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println(students)

	// The older style you will still see in other people's answers.
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] > numbers[j] })
	fmt.Println(numbers)
}

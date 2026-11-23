// Implicit interfaces, small interfaces, io.Reader, type switch, the nil trap.
//
//	go run ./05-interface/shapes
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type Shape interface {
	Area() float64
}

type Square struct{ Side float64 }
type Circle struct{ R float64 }

func (s Square) Area() float64 { return s.Side * s.Side }
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

// String satisfies fmt.Stringer, so Println uses this text.
func (s Square) String() string { return fmt.Sprintf("Square %.0f", s.Side) }

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// countLines accepts an io.Reader, so it works with files, stdin, or strings.
func countLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	n := 0
	for sc.Scan() {
		n++
	}
	return n, sc.Err()
}

func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("number %d", x)
	case string:
		return fmt.Sprintf("text %q", x)
	case Shape:
		return fmt.Sprintf("shape with area %.2f", x.Area())
	default:
		return fmt.Sprintf("other type %T", x)
	}
}

type CustomError struct{}

func (*CustomError) Error() string { return "custom" }

func wrong() error {
	var e *CustomError // nil
	return e           // this error interface is NOT nil
}

func right() error { return nil }

func main() {
	fmt.Printf("%.2f\n", totalArea([]Shape{Square{2}, Circle{1}}))
	fmt.Println(Square{3})

	n, _ := countLines(strings.NewReader("a\nb\nc"))
	fmt.Println("lines:", n)

	for _, v := range []any{7, "go", Circle{2}, 3.5} {
		fmt.Println(describe(v))
	}

	fmt.Println("wrong() != nil:", wrong() != nil)
	fmt.Println("right() != nil:", right() != nil)
}

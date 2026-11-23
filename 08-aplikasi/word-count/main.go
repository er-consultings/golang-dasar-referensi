// Command word-count counts lines, words, and the most frequent words
// in a file or in standard input.
//
//	go run ./08-aplikasi/word-count README.md
//	echo "one two two three three three" | go run ./08-aplikasi/word-count -top 2
package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"unicode"
)

// Stats is the result of counting a text.
type Stats struct {
	Lines int
	Words int
	Freq  map[string]int
}

// Count reads r until the end. Upper and lower case are treated the same,
// and punctuation at the start and end of a word is dropped.
func Count(r io.Reader) (Stats, error) {
	s := Stats{Freq: map[string]int{}}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		s.Lines++
		for _, w := range strings.Fields(sc.Text()) {
			w = strings.ToLower(strings.TrimFunc(w, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsDigit(r)
			}))
			if w == "" {
				continue
			}
			s.Words++
			s.Freq[w]++
		}
	}
	return s, sc.Err()
}

type wordCount struct {
	word  string
	count int
}

// Top returns the n most frequent words; ties are sorted alphabetically.
func (s Stats) Top(n int) []wordCount {
	out := make([]wordCount, 0, len(s.Freq))
	for w, c := range s.Freq {
		out = append(out, wordCount{w, c})
	}
	slices.SortFunc(out, func(a, b wordCount) int {
		if c := cmp.Compare(b.count, a.count); c != 0 {
			return c
		}
		return cmp.Compare(a.word, b.word)
	})
	return out[:min(n, len(out))]
}

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	fs := flag.NewFlagSet("word-count", flag.ContinueOnError)
	top := fs.Int("top", 5, "number of top words to show")
	if err := fs.Parse(args); err != nil {
		return err
	}

	in := stdin
	if fs.NArg() > 0 {
		f, err := os.Open(fs.Arg(0))
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	s, err := Count(in)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}
	fmt.Fprintf(stdout, "lines: %d, words: %d, unique words: %d\n", s.Lines, s.Words, len(s.Freq))
	for _, wc := range s.Top(*top) {
		fmt.Fprintf(stdout, "%-15s %d\n", wc.word, wc.count)
	}
	return nil
}

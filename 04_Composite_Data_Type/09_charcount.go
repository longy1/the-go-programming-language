package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	utflen := [utf8.UTFMax + 1]int{}
	invalid := 0

	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			if _, err := fmt.Fprintf(os.Stderr, "charcount: %v\n", err); err != nil {
			}
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, c := range counts {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("len\tcount\n")
	for l, c := range utflen {
		if l > 0 {
			fmt.Printf("%d\t%d\n", l, c)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

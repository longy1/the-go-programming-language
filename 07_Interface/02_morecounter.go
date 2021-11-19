package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))

	for scan.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	input := "The String method is used to print values passed\n" +
		"as an operand to any format that accepts a string\n" +
		"or to an unformatted printer such as Print."
	var lc LineCounter
	var wc WordCounter
	if _, err := fmt.Fprintf(&lc, input); err != nil {
	}
	if _, err := fmt.Fprintf(&wc, input); err != nil {
	}
	fmt.Println(lc, wc)
}

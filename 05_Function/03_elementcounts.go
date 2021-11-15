package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "element counts: %v\n", err); err != nil {
		}
	}

	counts := make(map[string]int)
	elementCount(counts, doc)
	fmt.Println("element counts:")
	for c, n := range counts {
		fmt.Printf("%-12.12s %d\n", c, n)
	}
}

func elementCount(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts["<"+n.Data+">"]++
	}
	if n.FirstChild != nil {
		elementCount(counts, n.FirstChild)
	}
	if n.NextSibling != nil {
		elementCount(counts, n.NextSibling)
	}
}

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func outTextContent(n *html.Node) {
	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) != "" {
			fmt.Println(n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n.Data != "script" && n.Data != "style" && n.Data != "textarea" {
			outTextContent(c)
		}
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		if err != nil {
		}
		os.Exit(1)
	}
	outTextContent(doc)
}

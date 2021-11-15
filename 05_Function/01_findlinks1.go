package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err); err != nil {
		}
		os.Exit(1)
	}

	f1, _ := os.OpenFile("./f1.txt", os.O_RDWR|os.O_CREATE, os.FileMode(0666))
	defer func(f1 *os.File) {
		err := f1.Close()
		if err != nil {
		}
	}(f1)
	for _, link := range visit(nil, doc) {
		if _, err := fmt.Fprintln(f1, link); err != nil {
		}
	}

	f2, _ := os.OpenFile("./f2.txt", os.O_RDWR|os.O_CREATE, os.FileMode(0666))
	defer func(f2 *os.File) {
		err := f2.Close()
		if err != nil {
		}
	}(f2)
	for _, link := range visitUsingLoopForChild(nil, doc) {
		if _, err := fmt.Fprintln(f2, link); err != nil {
		}
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func visitUsingLoopForChild(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.NextSibling != nil {
		links = visitUsingLoopForChild(links, n.NextSibling)
	}
	if n.FirstChild != nil {
		links = visitUsingLoopForChild(links, n.FirstChild)
	}
	return links
}

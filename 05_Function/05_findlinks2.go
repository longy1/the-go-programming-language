package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(links, c)
	}
	return links
}

func findlinks2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit2(nil, doc), nil
}

func main() {
	for _, arg := range os.Args[1:] {
		links, err := findlinks2(arg)
		if err != nil {
			if _, err := fmt.Fprintf(os.Stderr, "findlinks2: %v", err); err != nil {
			}
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

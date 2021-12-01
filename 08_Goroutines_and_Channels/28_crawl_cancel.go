// controlled concurrent spider, using bounded resources
// press enter to cancel
package main

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func crawl5(url string, ctx context.Context) []string {
	fmt.Println(url)
	list, err := Extract(url, ctx)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	var n int64
	n++
	go func() { worklist <- os.Args[1:] }()

	// feed unseenLinks
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl5(link, ctx)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	go func() {
		if _, err := os.Stdin.Read(make([]byte, 1)); err != nil {
		}
		cancel()
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				n++
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func Extract(url string, ctx context.Context) ([]string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		if err := resp.Body.Close(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("extract url %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	ForEachNode(doc, visitNode, nil)
	return links, nil
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

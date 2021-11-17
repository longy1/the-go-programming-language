package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
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

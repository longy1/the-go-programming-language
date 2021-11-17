package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

var depth int

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		ok := pre(n)
		if !ok {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		ok := post(n)
		if !ok {
			return n
		}
	}
	return nil
}

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		// handle attr
		var attrList []string
		var tail rune
		for _, attr := range n.Attr {
			attrList = append(attrList, fmt.Sprintf("%s=\"%s\"", attr.Key, attr.Val))
		}
		if n.FirstChild == nil {
			tail = '/'
		}
		fmt.Printf("%*s<%s %s%c>\n",
			depth*2, " ", n.Data, strings.Join(attrList, " "), tail)
		depth++
	}
	return true
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
		}
	}
	return true
}

func outline2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return err
	}
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func main() {
	for _, url := range os.Args[1:] {
		err := outline2(url)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
			}
		}
	}
}

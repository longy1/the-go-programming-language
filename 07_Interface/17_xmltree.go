// xmltree read xml doc from os.Stdin, and build a xml node tree
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

var depth int

func main() {
	doc := xml.NewDecoder(os.Stdin)
	var stack []Element
	for {
		tok, err := doc.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			child := Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) > 0 {
				stack[len(stack)-1].Children =
					append(stack[len(stack)-1].Children, child)
			}
			stack = append(stack, child)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			if len(stack) > 0 && len(tok) > 0 {
				stack[len(stack)-1].Children =
					append(stack[len(stack)-1].Children, CharData(tok))
			}
		}
	}
	if len(stack) <= 0 {
		return
	}
	visit(stack[0])
}

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func visit(n Node) {
	switch n := n.(type) {
	case CharData:
		fmt.Printf("%*s\n", depth*2, n)
	case Element:
		fmt.Printf("%*s<%s", depth*2, "", n.Type.Local)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Name.Local, attr.Value)
		}
		fmt.Printf(">\n")
		depth++
		for _, child := range n.Children {
			visit(child)
		}
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Type.Local)
	default:
		panic("unknown node")
	}
}

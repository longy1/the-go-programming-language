// controlled concurrent spider, using bounded resources
package main

import (
	"The.Go.Programming.Language/localpkg/links"
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	// avoid deadlock caused by unbuffered goroutine
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

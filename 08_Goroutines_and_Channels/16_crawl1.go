// uncontrolled concurrent spider, using too many resources
package main

import (
	"The.Go.Programming.Language/localpkg/links"
	"fmt"
	"log"
	"os"
)

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// avoid deadlock caused by unbuffered goroutine
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl1(link)
				}(link)
			}
		}
	}
}

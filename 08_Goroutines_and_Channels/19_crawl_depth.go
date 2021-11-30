// crawl optional depth links
package main

import (
	"The.Go.Programming.Language/localpkg/links"
	"flag"
	"fmt"
	"log"
)

var depth = flag.Int("d", 3, "limited crawl depth")

type linkItem struct {
	link  string
	depth int
}

func crawl4(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	flag.Parse()
	worklist := make(chan []linkItem)
	unseenLinks := make(chan linkItem)

	if len(flag.Args()) <= 0 {
		log.Fatal("Expected at least one url, but no url given.")
	}

	go func() {
		for _, link := range flag.Args() {
			unseenLinks <- linkItem{link, 0}
		}
	}()

	// feed unseenLinks
	for i := 0; i < 20; i++ {
		go func() {
			for it := range unseenLinks {
				if it.depth <= *depth {
					var foundLinks []linkItem
					for _, foundLink := range crawl4(it.link) {
						foundLinks = append(foundLinks,
							linkItem{foundLink, it.depth + 1})
					}
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, it := range list {
			if !seen[it.link] {
				seen[it.link] = true
				unseenLinks <- it
			}
		}
	}
}

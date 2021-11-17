package main

import (
	"The.Go.Programming.Language/localpkg/links"
	"fmt"
	"log"
	"os"
)

var originHost map[string]bool

// breadth first means visit next level after visiting current level
func breadthFirst(f func(item string) []string, workList []string) {
	for _, url := range workList {
		originHost[url] = true
	}
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	workList, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return workList
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

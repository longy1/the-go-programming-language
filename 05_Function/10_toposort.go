package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"discrete math"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming", "calculus"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	var visitAll func(items []string)
	seen := make(map[string]uint8)

	visitAll = func(items []string) {
		for _, item := range items {
			if seen[item] == 0 {
				seen[item] = 1
				visitAll(m[item])
				order = append(order, item)
				seen[item] = 2
			} else if seen[item] == 1 {
				fmt.Println("find cycle start with: " + item)
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

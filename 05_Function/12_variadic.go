package main

import (
	"fmt"
	"math"
	"strings"
)

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func max(vals ...int) int {
	m := -math.MaxInt64
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func min(vals ...int) int {
	m := math.MaxInt64
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

func Join(seq string, strs ...string) string {
	return strings.Join(strs, seq)
}

func main() {
	fmt.Println(sum(1, 3, 2, 4, 5))
	fmt.Println(sum())

	fmt.Println(max())
	fmt.Println(max(3, -2, 1, 8))

	fmt.Println(min())
	fmt.Println(min(3, -2, 1, 8))

	fmt.Println(Join(", ", "a", "b", "c", "haha"))
}

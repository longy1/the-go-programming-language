package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

func main() {
	f := square
	fmt.Println("square(3) =", f(3))

	f = negative
	fmt.Println("negative(5) =", f(5))

	g := product
	fmt.Println("product(3, 4) =", g(3, 4))

	fmt.Println(strings.Map(add1, "123456"))
	fmt.Println(strings.Map(add1, "abide"))
	fmt.Println(strings.Map(add1, "dijkstra"))
}

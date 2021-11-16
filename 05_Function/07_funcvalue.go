package main

import "fmt"

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	f := square
	fmt.Println("square(3) =", f(3))

	f = negative
	fmt.Println("negative(5) =", f(5))

	g := product
	fmt.Println("product(3, 4) =", g(3, 4))
}

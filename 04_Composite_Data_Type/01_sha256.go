package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("xx"))
	c2 := sha256.Sum256([]byte("XX"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	a := [4]int{4, 3, 2, 1}
	zero(&a)
	fmt.Println(a)

	fmt.Println(diffCountSha256(&c1, &c2))
}

func zero(a *[4]int) {
	for i := range a {
		a[i] = i
	}
}

func diffCountSha256(c1, c2 *[32]byte) int {
	count := 0
	for i := range c1 {
		xor := c1[i] ^ c2[i]
		for xor != 0 {
			if xor&1 == 1 {
				count++
			}
			xor >>= 1
		}
	}
	return count
}

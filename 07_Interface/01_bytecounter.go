package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	if _, err := c.Write([]byte("hello")); err != nil {
	}
	fmt.Println(c)

	c = 0
	name := "Dolly"
	if _, err := fmt.Fprintf(&c, "hello, %s", name); err != nil {
	}
	fmt.Println(c)
}

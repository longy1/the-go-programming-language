package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer f(%d)\n", x)
	f(x - 1)
}

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	_, err := os.Stdout.Write(buf[:n])
	if err != nil {
		return
	}
}

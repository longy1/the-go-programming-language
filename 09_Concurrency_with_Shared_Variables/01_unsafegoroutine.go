package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	var sg sync.WaitGroup
	sg.Add(3)
	go add(&sg)
	go add(&sg)
	go add(&sg)
	sg.Wait()
	fmt.Println(x)
}

func add(sg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		x++
	}
	sg.Done()
}

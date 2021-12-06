package main

import (
	"fmt"
	"time"
)

func main() {
	var head, now, next chan string
	head = make(chan string)
	now = head
	var count int
	for i := 0; i < 1000000; i++ {
		count++
		fmt.Println(count)
		next = make(chan string)
		go func(out, in chan string) {
			for msg := range in {
				out <- msg
			}
		}(next, now)
		now = next
	}
	head <- "message"
	start := time.Now()
	<-now
	fmt.Printf("%.4f\n", time.Since(start).Seconds())
}

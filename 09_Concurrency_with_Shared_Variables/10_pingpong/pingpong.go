package main

import (
	"fmt"
	"time"
)

var tick = time.Tick(1 * time.Second)

func main() {
	start := time.Now()
	ping := make(chan int)
	pong := make(chan int)
	go pingpong(ping, pong, start)
	go pingpong(pong, ping, start)
	ping <- 0
	for {
	}
}

func pingpong(out, in chan int, t time.Time) {
	for {
		select {
		case <-tick:
			count := <-in
			d := time.Since(t).Seconds()
			fmt.Printf("%.0f: %d avg: %.2f\n", d, count, float64(count)/d)
			out <- count + 1
		case count := <-in:
			out <- count + 1
		}
	}
}

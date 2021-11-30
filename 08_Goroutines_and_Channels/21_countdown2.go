package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		if _, err := os.Stdin.Read(make([]byte, 1)); err != nil {
			log.Println(err)
		}
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	// launch
	fmt.Println("Lift off!")
}

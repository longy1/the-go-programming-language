package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	// trace(msg) will be executed here, and defer it's return function
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Println("enter: " + msg)
	return func() {
		log.Printf("exit: %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}

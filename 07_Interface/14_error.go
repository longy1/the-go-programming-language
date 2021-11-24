package main

import (
	"fmt"
	"syscall"
)

func main() {
	for no := 0; no < 300; no++ {
		str := syscall.Errno(no).Error()
		fmt.Println(no, ":", str)
	}
}

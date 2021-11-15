// print os.Args[0]
// print index and values per line
// compare time using between strings.Join() and string add operation
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}

	var stringArray [10000]string
	for i := 0; i < 10000; i++ {
		stringArray[i] = string(rune(i))
	}
	s := ""
	start := time.Now()
	for i := 0; i < 10000; i++ {
		s += " " + stringArray[i]
	}
	cost1 := time.Since(start)

	s = ""
	start = time.Now()
	s = strings.Join(stringArray[:], " ")
	cost2 := time.Since(start)

	fmt.Println(cost1, cost2)
}

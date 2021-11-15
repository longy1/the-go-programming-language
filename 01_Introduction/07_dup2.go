// count repeat line in file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				if err != nil {
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	fmt.Println(f.Name())
	for input.Scan() {
		counts[input.Text()]++
	}
}

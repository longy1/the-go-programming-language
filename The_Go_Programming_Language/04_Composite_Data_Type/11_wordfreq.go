// workfreq counts the frequency of every word in input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCounts[input.Text()]++
	}
	fmt.Println("word\tcount")
	for w, n := range wordCounts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}

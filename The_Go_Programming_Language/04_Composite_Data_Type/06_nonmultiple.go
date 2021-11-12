package main

import "fmt"

func main() {
	strs := []string{"1", "1", "1", "2", "1", "3", "3", "33", "333", "33"}
	fmt.Println(strs)
	strs = nonMultiple(strs)
	fmt.Println(strs)
}

// no err dealing
func nonMultiple(strs []string) []string {
	out := strs[:0]
	for _, v := range strs {
		if len(out) == 0 || out[len(out)-1] != v {
			out = append(out, v)
		}
	}
	return out
}

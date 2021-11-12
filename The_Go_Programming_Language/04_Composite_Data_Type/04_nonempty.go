package main

import "fmt"

func main() {
	s1 := nonempty([]string{"a", "", "b", "c", ""})
	fmt.Println(s1)
	s2 := nonempty2([]string{"a", "", "b", "c", ""})
	fmt.Println(s2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename1("a"))
	fmt.Println(basename1("a.go"))
	fmt.Println(basename1("/a/b/c.py"))
	fmt.Println(basename1("/你/我/它.py"))

	fmt.Println(basename2("a"))
	fmt.Println(basename2("a.go"))
	fmt.Println(basename2("/a/b/c.py"))
	fmt.Println(basename2("/你/我/它.py"))
}

func basename1(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // return -1 if not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

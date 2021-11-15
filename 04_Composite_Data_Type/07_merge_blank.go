package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("\t\n  abcd \xff\xff \naa\t  bb\ncc")
	fmt.Println(string(s))
	s = mergeBlank(s)
	fmt.Println(string(s))
}

func mergeBlank(s []byte) []byte {
	head := 0
	afterBlank := false
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRune(s[i:])
		if r != utf8.RuneError {
			if unicode.IsSpace(r) {
				if afterBlank == true {
					i++
					continue
				}
				afterBlank = true
				copy(s[head:head+1], " ")
				i += 1
			} else {
				afterBlank = false
				copy(s[head:head+n], s[i:i+n])
				head += n
				i += n
			}
		} else {
			i++
		}
	}
	return s[:head]
}

package main

import "fmt"

func main() {
	s := [...]int{5, 4, 3, 2, 1}
	reverse(s[2:])
	reverse(s[:2])
	reverse(s[:])
	fmt.Println(s)
	reverseByPointer(&s)
	fmt.Println(s)
	rotate(s[:], 3)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseByPointer(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, dis int) {
	for i := range s {
		if i < dis {
			v := s[0]
			copy(s[:len(s)-1], s[1:])
			s[len(s)-1] = v
		}
	}
}

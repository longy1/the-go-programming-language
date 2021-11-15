// stack implemented by slice
package main

import "fmt"

func push(stack []int, value int) []int {
	return append(stack, value)
}

func pop(stack []int) ([]int, int) {
	if len(stack) <= 0 {
		return nil, -1
	}
	return stack[0 : len(stack)-1], stack[len(stack)-1]
}

func remove(stack []int, i int) []int {
	copy(stack[i:], stack[i+1:])
	return stack[0 : len(stack)-1]
}

func main() {
	var stack []int
	for i := 0; i < 10; i++ {
		stack = push(stack, i)
	}
	fmt.Println(stack)
	for i := 0; i < 10; i = i + 3 {
		var value int
		stack, value = pop(stack)
		fmt.Println(value)
	}
	fmt.Println(stack)
	stack = remove(stack, 1)
	stack = remove(stack, 1)
	fmt.Println(stack)
}

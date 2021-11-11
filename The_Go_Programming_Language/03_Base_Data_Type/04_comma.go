// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123345567"))
	fmt.Println(commaNonRecursive("123345567"))
	fmt.Println(commaNonRecursive("1123345567"))
	fmt.Println(commaNonRecursive("12"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNonRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	buf := bytes.Buffer{}
	r := n % 3
	if r != 0 {
		if _, err := fmt.Fprintf(&buf, "%s", s[:r]); err != nil {
		}
		buf.WriteByte(',')
	}
	for ; r < n; r += 3 {
		if _, err := fmt.Fprintf(&buf, "%s", s[r:r+3]); err != nil {
		}
		if r < n-3 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

// intsToString is like fmt.Sprint(values) but adds commas.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	values := []int{1, 3, 2, 4, 5, 3}
	fmt.Println(intsToString(values))
}

func intsToString(values []int) string {
	var buf = bytes.Buffer{}
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		if _, err := fmt.Fprintf(&buf, "%d", v); err != nil {
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

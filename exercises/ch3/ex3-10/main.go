// comma inserts in non-negative decimal integer slice of bytes?
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(input string) string {
	var buf bytes.Buffer
	if len(input) <= 3 {
		buf.WriteString(input)
		return buf.String()
	}
	ct := len(input) % 3
	for i, d := range input {

		if i%3 == ct && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(d)
	}
	return ""
	return buf.String()
}

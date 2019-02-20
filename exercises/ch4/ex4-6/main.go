package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	testBytes := []byte("界  hey    team . .  .   ")
	spaceCheck(testBytes)
	fmt.Println(string(testBytes))

}

func spaceCheck(bs []byte) []byte {
	c := 0
	for i := 0; i < len(bs); i++ {
		r, _ := utf8.DecodeRune(bs[i:])
		var prevSpace = false
		if unicode.IsSpace(r) {
			fmt.Println("found space", r, c, i)
			c++
			prevSpace = true
			continue
		}
		prevSpace = false
		if !prevSpace {
			fmt.Println(string(bs[c:c+1]), string(bs[i:i+1]))
			copy(bs[c:c+1], bs[i:i+1])
		}

	}
	return bs
}

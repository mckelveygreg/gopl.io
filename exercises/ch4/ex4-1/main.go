package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("Hey Team!"))
	c2 := sha256.Sum256([]byte("11"))

	// diff := differences(c1, c2)

	fmt.Println(c1, c2)
}

func differences(a, b [32]byte) int {
	var c int

	for i, v := range a {
		c += setBits(v ^ b[i])
	}
	return c
}

func setBits(b byte) int {
	var c int

	for b > 0 {
		b = b & (b - 1)
		c++
	}
	return c
}

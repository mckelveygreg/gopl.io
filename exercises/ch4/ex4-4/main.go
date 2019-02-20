// ex 4.4
// Write a version of rotate (pg. 92) that operates in a single pass

package main

import (
	"fmt"
)

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// whats the diff between these two arrays??
	b := []int{0, 1, 2, 3, 4, 5}

	// one liner rotate
	b = append(b[2:], b[:2]...)
	fmt.Println(b)
}


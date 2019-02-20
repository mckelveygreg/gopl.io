package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reversePointer(&a)

	fmt.Println(a)
}

// So, Go says not to pass array pointers and to use slices instead? What gives?
func reversePointer(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

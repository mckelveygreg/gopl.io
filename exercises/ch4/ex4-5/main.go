package main

import (
	"fmt"
)

func dedup(s []string) []string {
	var dd []string
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			dd = append(dd, s[i])
		}
	}
	return dd
}

func davedDup(s []string) []string {
	c := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[c-1] {
			s[c] = s[i]
			c++
		}
	}
	return s[:c:c]
}

func main() {
	slice := []string{"Hey", "Hey", "Hey", "You", "Team", "Team", "Captain", "Captain"}
	fmt.Println(dedup(slice))
	slice = []string{"All", "the", "the", "things"}
	fmt.Println(dedup(slice))
}

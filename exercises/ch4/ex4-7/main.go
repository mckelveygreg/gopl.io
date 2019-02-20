package main

import "fmt"

func main() {
	bs := []byte("racecar")
	reverse(bs)
	fmt.Println(string(bs))
	bs = []byte("th界is is a test")
	reverse(bs)
	fmt.Println(string(bs))
}

func reverse(bs []byte) []byte {
	var count = len(bs) / 2
	for i := 0; i < count; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	return bs
}

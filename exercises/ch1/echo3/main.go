// Echo 3 prints its command-line agruments, but faster
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(os.Args[0], "\n", strings.Join(os.Args[1:], " ")) // ex 1.1
	//fmt.Println(strings.Join(os.Args, " "))
}

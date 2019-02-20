// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	//s := ""
	//sep := ""
	for i, arg := range os.Args[1:] {
		//sep = "\n"
		fmt.Printf("%d %s\n", i, arg) //ex 1.2
		//s += sep + i + " " + arg
	}
	//fmt.Println(s)
}

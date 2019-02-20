// Dup2 prints the count and text of lines hat appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int) // dup-text {file: count}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			file := files[i]
			fmt.Println(file)
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fileNames := range counts {
		for fileName, ct := range fileNames {
			if ct > 1 {
				fmt.Printf("%s\t%s\t%d\n", line, fileName, ct)
			}
		}
	}
	fmt.Println(counts)
}
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}

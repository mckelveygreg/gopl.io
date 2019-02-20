// wordfreq reports the frequency of each word in a input text file
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	f, err := os.Open("hip.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		txt := strings.ToLower(sc.Text())
		txt = strings.TrimFunc(txt, func(r rune) bool { return unicode.IsPunct(r) })
		counts[txt]++
	}

	// fmt.Printf("word\t\tcount\n")

	// for w, c := range counts {
	// 	fmt.Printf("%s\t%d\n", w, c)
	// }

	//sortAlpha(counts)

	sortByCount(counts)
}

func sortAlpha(counts map[string]int) {
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, counts[k])
	}
}

func sortByCount(counts map[string]int) {
	valueMap := make(map[int][]string)
	for k, v := range counts {
		valueMap[v] = append(valueMap[v], k)
	}
	values := make([]int, 0, len(counts))
	for v := range valueMap {
		values = append(values, v)
	}
	sort.Ints(values)
	for _, v := range values {
		fmt.Println(v, valueMap[v])
	}
}

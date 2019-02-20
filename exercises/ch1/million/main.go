// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "million: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\r\n") {
			go fetch(line, ch)
		}
	}
	// for _, url := range os.Args[1:] {
	// 	go fetch(url, ch) // start a goroutine
	// }
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot read file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds()) // how to modify and log new changes?
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

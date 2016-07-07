// Package main provides ...
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	var result string
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}
	for _, url := range os.Args[1:] {
		result += <-ch + "\n" // Receive from channel ch. Using same channel for all receive operations ensures that we print full string from one sender before proceeding with the next sender
		go fetch(url, ch)     // Send another request to check effect of caching. Sending a request at this point ensures that previous response was complete before we attempted second request. So, caching (if any) should be in full-effect by this point.
		result += <-ch + "\n" // For the second response
	}
	result += fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())

	ioutil.WriteFile("result.txt", []byte(result), 0644)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f  %7d  %s", secs, nbytes, url)
}

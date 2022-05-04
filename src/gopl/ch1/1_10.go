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
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch_two_sequence(url, ch)
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}

func fetch_two_sequence(url string, ch chan<- string) {
	fetch(url, ch)
	fetch(url, ch)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
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
	ch <- fmt.Sprintf("Status: %v %.2fs  %7d  %s", resp.Status, secs, nbytes, url)
}

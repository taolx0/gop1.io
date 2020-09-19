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
	start := time.Now()
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for _, _ = range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs escapled\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %v:%s\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %8d %s\n", secs, nBytes, url)
}

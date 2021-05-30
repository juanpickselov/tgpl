// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	chnl := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, chnl)
	}
	for range os.Args[1:] {
		fmt.Println(<-chnl)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, chnl chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		chnl <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		chnl <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	chnl <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

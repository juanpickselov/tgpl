package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutx sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(wrtr http.ResponseWriter, rqust *http.Request) {
	mutx.Lock()
	count++
	mutx.Unlock()
	fmt.Fprintf(wrtr, "URL.Path = %q\n", rqust.URL.Path)
}

func counter(wrtr http.ResponseWriter, rqust *http.Request) {
	mutx.Lock()
	fmt.Fprintf(wrtr, "Count %d\n", count)
}

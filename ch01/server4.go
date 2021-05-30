package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(wrtr http.ResponseWriter, rqst *http.Request) {
	fmt.Fprintf(wrtr, "URL.Path = %q\n", rqst.URL.Path)
}

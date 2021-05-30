package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(wrtr http.ResponseWriter, rqust *http.Request) {
	fmt.Fprintf(wrtr, "%s %s %s\n", rqust.Method, rqust.URL, rqust.Proto)
	for theKey, theValue := range rqust.Header {
		fmt.Fprintf(wrtr, "Header[%q] = %q\n", theKey, theValue)
	}
	fmt.Fprintf(wrtr, "Host = %q\n", rqust.Host)
	fmt.Fprintf(wrtr, "RemoteAddr = %q\n", rqust.RemoteAddr)
	if err := rqust.ParseForm(); err != nil {
		log.Print(err)
	}
	for theKey, theValue := range rqust.Form {
		fmt.Fprintf(wrtr, "Form[%q] = %q\n", theKey, theValue)
	}
}

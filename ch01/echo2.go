// Echo2 prints its command-line args args args.package ch01
package main

import (
	"os"
	"fmt"
)

func main() {
	aString, separator := "", ""
	for _, arg := range os.Args[1:] {
		aString += separator + arg
		separator = " "
	}
	fmt.Println(aString)
}

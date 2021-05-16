// Echo2 prints its command-line args args args.package ch01
package main

import (
	"fmt"
	"os"
)

func main() {
	for indexNum, arg := range os.Args[1:] {
		fmt.Printf("%d "+arg+"\n", indexNum)
	}
}

// Echo4 isn't even in the book but I am doing it just because
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0] + " which was run with the args below")
	fmt.Println(os.Args[1:])
}

// Dup3 reads files showing duplicate line counts & text
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, numberLines := range counts {
		if numberLines > 1 {
			fmt.Printf("%d\t%s\n", numberLines, line)
		}
	}
}

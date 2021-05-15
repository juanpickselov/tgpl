// Dup1 prints count and text of each duplicate line from standard input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, numberLines := range counts {
		if numberLines > 1 {
			fmt.Printf("%d\t%s\n", numberLines, line)
		}
	}
}

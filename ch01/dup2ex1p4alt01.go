// Dup2 prints count, duplicate text lines, and filenames if used, from standard input or file list
// Solution gist used from: https://gist.github.com/limingjie/be429b2eb2c48795f32522337bfd5a1d
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			fileToCheck, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(fileToCheck, arg, counts)
			fileToCheck.Close()
		}
	}
	for line, filenames := range counts {
		for fName, ct := range filenames {
			if ct > 1 {
				fmt.Println(line, "is duplicated in", fName, ct, "times")
			}
		}
	}
}

func countLines(fileToCheck *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(fileToCheck)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}

// Dup2 prints count and text of duplicate lines from standard input or file list
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileList := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileList)
	} else {
		for _, arg := range files {
			fileToCheck, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(fileToCheck, counts, fileList)
			fileToCheck.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, fileList[line], line)
		}
	}
}

func in(needle string, strings []string) bool {
	//Do not like this looking at alternative solutions
	for _, aString := range strings {
		if needle == aString {
			return true
		}
	}
	return false
}

func countLines(fileToCheck *os.File, counts map[string]int, fileList map[string][]string) {
	input := bufio.NewScanner(fileToCheck)
	for input.Scan() {
		line := input.Text()
		counts[input.Text()]++
		if !in(fileToCheck.Name(), fileList[line]) {
			fileList[line] = append(fileList[line], fileToCheck.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

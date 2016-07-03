// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"os"
)

type countAndName struct {
	count int
	name  string
}

func main() {
	counts := make(map[string]*countAndName)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%s: %d in files: %s\n", line, n.count, n.name)
		}
	}
}

func countLines(f *os.File, counts map[string]*countAndName) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		if _, ok := counts[txt]; !ok {
			counts[txt] = new(countAndName)
		}
		counts[txt].count++
		if counts[txt].name != "" {
			counts[txt].name += ", "
		}
		counts[txt].name += f.Name()
	}
}

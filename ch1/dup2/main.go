// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// func Open(name string) (*File, error)

// stderr, stdin, stdout is just three different file.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, cnt map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		cnt[input.Text()]++
	}
}

func main() {
	cnt := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countLines(os.Stdin, cnt)
	} else {
		for _, name := range(file) {
			f, err := os.Open(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, cnt)
			f.Close()
		}
	}

	for line, n := range cnt {
		if n > 1 {
			fmt.Fprintf(os.Stderr, "%s %d\n", line, n)
		}
	}
}
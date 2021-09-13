// read the entire input into memory in one big gulp, split it into
// lines all at once, then pro cess the lines.

// strings.Split
// func ReadFile(filename string) ([]byte, error)
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	cnt := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			cnt[line]++
		}
	}
}

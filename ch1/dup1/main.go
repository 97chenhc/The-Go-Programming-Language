// map[type]type
// control + D: EOF
/*
bufio: to read standard input as a set of lines
func (s *Scanner) Text() string
func NewScanner(r io.Reader) *Scanner
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cnt[input.Text()]++
	}
	for key, value := range cnt {
		fmt.Printf("%s %d\n", key, value)
	}
}
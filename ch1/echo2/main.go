// print the command line arguments
// use range for
package main

import (
	"os"
	"fmt"
)

func main() {
	s := ""
	// return index and element
	for _, arg := range os.Args[1:] {
		s += arg;
	}
	fmt.Println(s)
}
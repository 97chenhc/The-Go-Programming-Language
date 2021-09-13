// print the command line arguments
package main

import (
	"os"
	"fmt"
)

func main() {
	s := ""
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
	}
	fmt.Println(s)
}
// Echo4 prints its command-line arguments.
// flag
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(os.Args[1:], *sep))
	if !*n {
		fmt.Println()
	}
}
// io.Copy(dst, src)
// Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
// buffer large enough to hold the entire stream.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
	}
}
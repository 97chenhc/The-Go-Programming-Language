package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	var a tempconv.Celsius = 1.0
	b := tempconv.CToF(a)
	fmt.Printf("%g", b);
}
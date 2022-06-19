package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Printf("\n n pointer %g n value %g sep pinter %g sep value %g", n, *n, sep, *sep)
	if !*n {
		fmt.Println()
	}
}

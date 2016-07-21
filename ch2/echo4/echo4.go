package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newling")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

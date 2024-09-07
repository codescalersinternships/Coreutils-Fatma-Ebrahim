package main

import (
	
	"flag"
	"fmt"
	"strings"
	
)

func main() {

	nptr := flag.Bool("n", false, "omit echoing trailing newline")
	flag.Parse()
	input := strings.Join(flag.Args(), " ")
	
	n := *nptr
	if n {
		fmt.Print( input)

	} else{
		fmt.Println(input)
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	var repeat string

	if len(flag.Args()) == 1 {
		repeat = flag.Args()[0]
	} else {
		repeat = "y"
	}

	for {
		fmt.Println(repeat)
	}

}

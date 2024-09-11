package main

import (
	"flag"
	"fmt"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tail"
)

func main() {
	n := 10
	flag.IntVar(&n, "n", 10, "number of lines read")
	flag.Parse()
	filename := flag.Arg(0)
	err := tail.Tail(filename, n)
	if err != nil {
		fmt.Print(err)
	}

}

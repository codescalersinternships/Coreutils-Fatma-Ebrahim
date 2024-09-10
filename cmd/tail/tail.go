package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tail"
)

func main() {
	nptr := flag.Int("n", 10, "number of lines read")
	flag.Parse()
	filename := flag.Args()[0]
	n := *nptr
	tail.Tail(filename, n)

}

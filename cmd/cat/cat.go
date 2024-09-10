package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/cat"
)

func main() {

	nptr := flag.Bool("n", false, "number output lines")
	flag.Parse()
	filename := flag.Args()[0]
	n := *nptr

	cat.Cat(filename, n)
}

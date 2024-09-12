package main

import (
	"flag"
	"log"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/head"
)

func main() {
	var n uint
	flag.UintVar(&n, "n", 10, "number of lines read")
	flag.Parse()
	filename := flag.Arg(0)
	err := head.Head(filename, n)
	if err != nil {
		log.Fatal(err)
	}
}

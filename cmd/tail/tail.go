package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tail"
	"log"
)

func main() {
	var n uint
	flag.UintVar(&n, "n", 10, "number of lines read")
	flag.Parse()
	filename := flag.Arg(0)
	err := tail.Tail(filename, n)
	if err != nil {
		log.Fatal(err)
	}

}

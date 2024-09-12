package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tree"
	"log"
)

func main() {
	var l uint
	flag.UintVar(&l, "L", 1, "print files and directories up to 'num' levels of depth")
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	err := tree.Tree(dir, l)
	if err != nil {
		log.Fatal(err)
	}

}

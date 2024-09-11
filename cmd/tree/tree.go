package main

import (
	"flag"
	"fmt"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tree"
)

func main() {
	l := 1
	flag.IntVar(&l, "L", 1, "print files and directories up to 'num' levels of depth")
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	err:=tree.Tree(dir, l)
	if err != nil {
		fmt.Print(err)
	}

}

package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/tree"
)

func main() {
	lptr := flag.Int("L", 1, "print files and directories up to 'num' levels of depth")
	flag.Parse()
	var dir string

	if len(flag.Args()) == 1 {
		dir = flag.Args()[0]
	} else {
		dir = "."
	}
	l := *lptr
	tree.Tree(dir, l, l)

}

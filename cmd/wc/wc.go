package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/wc"
)

func main() {
	lptr := flag.Bool("l", false, "display number of lines only")
	wptr := flag.Bool("w", false, "display number of words only")
	cptr := flag.Bool("c", false, "display number of chars only")
	flag.Parse()
	filename := flag.Args()[0]
	wc.Wc(filename,lptr,wptr,cptr)

}

package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/wc"
	"log"
)

func main() {
	var l, w, c bool
	flag.BoolVar(&l, "l", false, "display number of lines only")
	flag.BoolVar(&w, "w", false, "display number of words only")
	flag.BoolVar(&c, "c", false, "display number of chars only")
	flag.Parse()
	filename := flag.Arg(0)
	err := wc.Wc(filename, l, w, c)
	if err != nil {
		log.Fatal(err)
	}

}

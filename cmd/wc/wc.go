package main

import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/wc"
)

func main() {
	l := false
	w := false
	c := false
	flag.BoolVar(&l, "l", false, "display number of lines only")
	flag.BoolVar(&w, "w", false, "display number of words only")
	flag.BoolVar(&c, "c", false, "display number of chars only")
	flag.Parse()
	var scanner *bufio.Scanner
	if len(flag.Args()) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
		
	}
	filename := flag.Arg(0)
	err:=wc.Wc(filename, l, w, c,scanner)
	if err != nil {
		fmt.Print(err)
	}

}

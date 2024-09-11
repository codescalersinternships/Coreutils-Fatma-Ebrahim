package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/cat"
)

func main() {
	n := false
	flag.BoolVar(&n, "n", false, "number output lines")
	flag.Parse()
	var scanner *bufio.Scanner
	if len(flag.Args()) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
	}

	filename := flag.Arg(0)
	err := cat.Cat(filename, n, scanner)
	if err != nil {
		fmt.Print(err)
	}
}

package main

import (
	"bufio"
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/cat"
	"log"
	"os"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "number output lines")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	if len(flag.Args()) != 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	}
	cat.Cat(n, scanner)

}

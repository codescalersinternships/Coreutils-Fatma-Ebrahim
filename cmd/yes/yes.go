package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/yes"
)

func main() {
	flag.Parse()
	repeat := ""

	if len(flag.Args()) == 1 {
		repeat = flag.Args()[0]
	}
	yes.Yes(repeat)

}

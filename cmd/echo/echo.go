package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/echo"
	"strings"
)

func main() {

	nptr := flag.Bool("n", false, "omit echoing trailing newline")
	flag.Parse()
	input := strings.Join(flag.Args(), " ")
	n := *nptr
	echo.Echo(input,n)
}

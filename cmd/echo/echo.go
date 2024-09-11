package main

import (
	"flag"
	"strings"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/echo"
)

func main() {
	n := false
	flag.BoolVar(&n, "n", false, "omit echoing trailing newline")
	flag.Parse()
	input := strings.Join(flag.Args(), " ")
	echo.Echo(input, n)
}

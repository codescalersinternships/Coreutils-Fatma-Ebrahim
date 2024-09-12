package main

import (
	"flag"
	"strings"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/echo"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "omit echoing trailing newline")
	flag.Parse()
	input := strings.Join(flag.Args(), " ")
	echo.Echo(input, n)
}

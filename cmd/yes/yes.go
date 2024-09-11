package main

import (
	"flag"
	"github.com/codescalersinternships/Coreutils-Fatma-Ebrahim/internal/yes"
)

func main() {
	flag.Parse()
	repeat:= flag.Arg(0)
	if repeat==""{
		repeat="y"
	}
	yes.Yes(repeat)

}

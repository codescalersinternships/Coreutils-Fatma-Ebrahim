package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lptr := flag.Bool("l", false, "display number of lines only")
	wptr := flag.Bool("w", false, "display number of words only")
	cptr := flag.Bool("c", false, "display number of chars only")
	flag.Parse()
	filename := flag.Args()[0]

	l := 0
	w := 0
	c := 0

	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		chars := strings.Split(line, "")
		w += len(words)
		c += len(chars)
		l++
	}

	switch {
	case *lptr:
		fmt.Print(l, " ",filename)
	case *wptr:
		fmt.Print(w, " ",filename)
	case *cptr:
		fmt.Print(c, " ",filename)
	default:
		fmt.Println(l, w, c,filename)

	}

	
	f.Close()

}

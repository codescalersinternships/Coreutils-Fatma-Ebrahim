package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	nptr := flag.Bool("n", false, "number output lines")
	flag.Parse()
	filename := flag.Args()[0]
	n := *nptr
	
	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)

	s := make([]string, 0)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if n {
		for j := 0; j < len(s); j++ {
			fmt.Println(" ",j+1," ",s[j])
		}

	} else {
		for j := 0; j < len(s); j++ {
			fmt.Println(s[j])
		}
	}

	f.Close()
}

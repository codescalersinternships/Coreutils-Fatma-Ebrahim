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

	nptr := flag.Int("n", 10, "number of lines read")
	flag.Parse()
	filename := flag.Args()[0]
	n := *nptr
	fmt.Println(n)

	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)

	s := make([]string, 0)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if n < len(s) {
		h := s[:n]
		for j := 0; j < len(h); j++ {
			fmt.Println(h[j])
		}

	} else {
		for j := 0; j < len(s); j++ {
			fmt.Println(s[j])
		}
	}

	f.Close()
}

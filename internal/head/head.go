package head

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Head(filename string, n int) {

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

package cat

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

func Cat(filename string, n bool) {

	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)

	s := make([]string, 0)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if n {
		for j := 0; j < len(s); j++ {
			fmt.Println(" ", j+1, " ", s[j])
		}

	} else {
		for j := 0; j < len(s); j++ {
			fmt.Println(s[j])
		}
	}

	f.Close()
}

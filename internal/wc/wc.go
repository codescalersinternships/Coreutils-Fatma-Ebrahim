package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Wc(filename string, lptr,wptr,cptr *bool) {


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

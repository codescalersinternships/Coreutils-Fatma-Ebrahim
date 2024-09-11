package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Wc(filename string, l, w, c bool, scanner *bufio.Scanner) error {

	linecount := 0
	wordcount := 0
	charcount := 0

	if scanner == nil {
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Can't open the file with path %q due error: %w", filename, err)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordcount += len(words)
		charcount += len(line) + 2
		linecount++
	}

	if !l && !w && !c {
		fmt.Println(linecount, wordcount, charcount, filename)

	} else {
		if l {
			fmt.Print(linecount, "  ")
		}
		if w {
			fmt.Print(wordcount, "  ")
		}
		if c {
			fmt.Print(charcount, "  ")
		}
		fmt.Print(filename)
	}

	return nil

}

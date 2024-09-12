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

		bytes, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("Can't open the file with path %q due error: %w", filename, err)
		}
		for _, e := range bytes {
			if string(e) == "\n" {
				linecount++
			}
			charcount += len(string(e))
		}
		words := strings.Fields(string(bytes))
		wordcount += len(words)
		

	} else {
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)
			wordcount += len(words)
			charcount += len(line) + 1
			linecount++
		}
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
		fmt.Println(filename)
	}

	return nil

}

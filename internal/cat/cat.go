package cat

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(filename string, n bool, scanner *bufio.Scanner) error {

	if scanner == nil {
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Can't open the file with path %q due error: %w", filename, err)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	}
	
	linenumber := 1
	for scanner.Scan() {
		if n {
			fmt.Print("    ", linenumber, "  ")
			linenumber++
		}
		fmt.Println(scanner.Text())
	}

	return nil

}

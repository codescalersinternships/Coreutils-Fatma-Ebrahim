package cat

import (
	"bufio"
	"fmt"
)

func Cat(n bool, scanner *bufio.Scanner)  {
	if scanner == nil {
		linenumber := 1
		for scanner.Scan() {
			if n {
				fmt.Print("    ", linenumber, "  ")
				linenumber++
			}
			fmt.Println(scanner.Text())
		}

	} else {
		linenumber := 1
		for scanner.Scan() {
			if n {
				fmt.Print("    ", linenumber, "  ")
				linenumber++
			}
			fmt.Println(scanner.Text())

		}
	}

}

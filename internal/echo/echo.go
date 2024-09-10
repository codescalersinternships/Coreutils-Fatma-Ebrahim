package echo

import (
	"fmt"
)

func Echo(input string, n bool) {

	if n {
		fmt.Print(input)

	} else {
		fmt.Println(input)
	}
}

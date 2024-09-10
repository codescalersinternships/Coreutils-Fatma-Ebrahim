package yes

import (
	"fmt"
)

func Yes(repeat string) {

	if repeat == "" {
		repeat = "y"
	}

	for {
		fmt.Println(repeat)
	}

}

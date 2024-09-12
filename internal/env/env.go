package env

import (
	"fmt"
	"os"
)

func Env() {

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

}

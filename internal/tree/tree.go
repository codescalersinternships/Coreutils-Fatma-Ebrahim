package tree

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Tree(dir string, level, depth int) {
	if level == 0 {
		return
	}
	level--
	c, err := os.ReadDir(dir)
	check(err)

	for _, entry := range c {
		fmt.Println(strings.Repeat("-----", depth-level), entry.Name())
		if entry.IsDir() {
			Tree(dir+"/"+entry.Name(), level, depth)
		}

	}
}

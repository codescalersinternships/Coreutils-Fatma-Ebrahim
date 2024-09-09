package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lptr:=flag.Int("L",1,"print files and directories up to 'num' levels of depth")
	flag.Parse()
	var dir string

	if len(flag.Args()) == 1 {
		dir = flag.Args()[0]
	} else {
		dir = "."
	}
	l:=*lptr
	treerecursive(dir, l, l)

}

func treerecursive(dir string, level, depth int) {
	if level == 0 {
		return
	}
	level--
	c, err := os.ReadDir(dir)
	check(err)

	for _, entry := range c {
		fmt.Println(strings.Repeat("-----", depth-level), entry.Name())
		if entry.IsDir() {
			treerecursive(dir+"/"+entry.Name(), level, depth)
		}

	}
}

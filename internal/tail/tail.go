package tail

import (
	"bufio"
	"fmt"
	"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Tail(filename string, n int) {

	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)
	
	s := make([]string, 0)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}


	if n<len(s){
		t:=s[len(s)-n:]
		for j:=0;j<len(t);j++{
			fmt.Println(t[j])
		}
	
	}else{
		for j:=0;j<len(s);j++{
			fmt.Println(s[j])
		}
	}


	
	f.Close()
}

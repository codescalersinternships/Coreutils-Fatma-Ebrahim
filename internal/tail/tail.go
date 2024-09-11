package tail

import (
	"bufio"
	"fmt"
	"os"
)

func Tail(filename string, n int) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Can't open the file with path %q due error: %w", filename, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if n > len(lines) {
		n = len(lines)
	}
	for j := len(lines) - n; j < len(lines); j++ {
		fmt.Println(lines[j])
	}
	return nil

}

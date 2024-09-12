package tail

import (
	"bufio"
	"fmt"
	"os"
)

func Tail(filename string, n uint) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Can't open the file with path %q due to error: %w", filename, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if n > uint(len(lines)) {
		n = uint(len(lines))
	}
	for j := uint(len(lines)) - n; j < uint(len(lines)); j++ {
		fmt.Println(lines[j])
	}
	return nil

}

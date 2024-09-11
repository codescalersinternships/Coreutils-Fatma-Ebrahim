package head

import (
	"bufio"
	"fmt"
	"os"
)

func Head(filename string, n int) error {

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Can't open the file with path %q due error: %w", filename, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() && n != 0 {
		fmt.Println(scanner.Text())
		n--
	}
	return nil

}

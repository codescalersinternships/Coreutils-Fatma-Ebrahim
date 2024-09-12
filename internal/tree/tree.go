package tree

import (
	"fmt"
	"os"
)

func Tree(dir string, level uint) error {
	var dirscount int
	var filescount int
	err := tree(dir, level, level, "", &dirscount, &filescount)
	if err != nil {
		return err
	}
	fmt.Printf("%d directories, %d files", dirscount, filescount)
	return nil
}
func tree(dir string, level, depth uint, prefix string, dirscount, filescount *int) error {
	if level == 0 {
		return nil
	}
	level--
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Can't Read directory with path %q due to error: %w", dir, err)
	}

	for index, entry := range entries {
		isLast := index == len(entries)-1
		branch := "├── "
		if isLast {
			branch = "└── "
		}

		if entry.IsDir() {
			*dirscount++
			fmt.Println(prefix + branch + entry.Name())
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			_ = tree(dir+"/"+entry.Name(), level, depth, newPrefix, dirscount, filescount)
		} else {
			*filescount++
			fmt.Println(prefix + branch + entry.Name())
		}
	}
	return nil
}

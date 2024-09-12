# CoreUtils in Go

This repository contains simple implementations of some core Unix utilities:

1. **head**: Prints the first 10 lines of a file by default.  
   - Option: `-n` to specify the number of lines to print.

2. **tail**: Prints the last 10 lines of a file by default.  
   - Option: `-n` to specify the number of lines to print.

3. **wc**: Counts the number of lines, words, and characters in a file.  
   - Options:  
     - `-l` for lines only  
     - `-w` for words only  
     - `-c` for characters only

4. **cat**: Concatenates and prints file contents.  
   - Option: `-n` to number the output lines.

5. **echo**: Prints the provided arguments to the standard output.  
   - Option: `-n` to omit the trailing newline.

7. **tree**: Displays the directory structure in a tree-like format.  
   - Options:  
     - `-L` to limit the display depth of the directory tree

6. **env**: Prints the environment variables.  

8. **true**: Always returns an exit status of 0 (success).  

9. **false**: Always returns an exit status of 1 (failure).  

10. **yes**: Repeatedly prints a specified string, or `y` by default.  

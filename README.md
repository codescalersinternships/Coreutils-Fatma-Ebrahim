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

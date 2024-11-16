# go-ls

This Go program emulates the basic functionality of the Unix `ls` command by listing and categorizing the contents of the current working directory into directories, files, and special files.

## Features

- **Directory Listing**: Displays all entries in the current working directory.
- **Categorization**: Separates entries into directories, regular files, and special files (e.g., symbolic links, sockets).
- **Numbered Output**: Provides a numbered list for each category.

## Usage

1. **Clone the Repository**: Ensure you have [Go installed](https://golang.org/dl/) on your system.

2. **Build the Program**:

   ```bash
   go build -o go-ls
   ```

3. **Run the Program**:

   ```bash
   ./go-ls
   ```

   The program will output the contents of the current directory, categorized as follows:

   ```
   Directories:
   1 - dir1
   2 - dir2

   Files:
   1 - file1.txt
   2 - file2.go

   Special Files:
   1 - symlink1
   ```

#
## Dependencies

This program relies solely on Go's standard library, specifically the `os`, `io/fs`, `log`, and `fmt` packages.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Inspired by the Unix `ls` command, this project serves as a learning tool for understanding directory operations in Go. 
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't determine current working directory:", err)
	}

	entries, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	totalEntries := len(entries)

	files := make([]fs.DirEntry, 0, totalEntries)
	directories := make([]fs.DirEntry, 0, totalEntries)
	special := make([]fs.DirEntry, 0, 0)

	for _, entry := range entries {
		switch mode := entry.Type(); {
		case mode.IsRegular():
			files = append(files, entry)
		case mode.IsDir():
			directories = append(directories, entry)
		default:
			special = append(special, entry)
		}
	}

	var output strings.Builder

	if len(directories) > 0 {
		output.WriteString("Directories:\n---\n")
		for i, dir := range directories {
			output.WriteString(fmt.Sprintf("%d - %v\n", i+1, dir.Name()))
		}
	}

	if len(files) > 0 {
		output.WriteString("\nFiles:\n---\n")
		for i, file := range files {
			output.WriteString(fmt.Sprintf("%d - %v\n", i+1, file.Name()))
		}
	}

	if len(special) > 0 {
		output.WriteString("\nSpecial Files:\n---\n")
		for i, sp := range special {
			output.WriteString(fmt.Sprintf("%d - %v\n", i+1, sp.Name()))
		}
	}

	fmt.Printf("%s\n", output.String())
}

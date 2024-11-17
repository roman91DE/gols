package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	hidden := flag.Bool("a", false, "Include hidden Files - Default is False")
	flag.Parse()

	inputDir := flag.Arg(0)
	var err error

	if inputDir == "" || inputDir == "." {
		inputDir, err = os.Getwd()
		if err != nil {
			log.Fatal("Couldn't determine current working directory:", err)
		}
	}

	if !pathExists(inputDir) {
		log.Fatal(fmt.Sprintf("Directory %s doesn't exists\n", inputDir))
	}

	output, err := goLs(inputDir, *hidden)
	if err != nil {
		log.Fatal("Error running go-ls!\n")
	}

	fmt.Printf("%s\n", output)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func goLs(dir string, hidden bool) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	totalEntries := len(entries)

	files := make([]fs.DirEntry, 0, totalEntries)
	directories := make([]fs.DirEntry, 0, totalEntries)
	special := make([]fs.DirEntry, 0, 0)

	hiddenRegex := regexp.MustCompile(`^\.`)

	for _, entry := range entries {
		if !hidden && hiddenRegex.MatchString(entry.Name()) {
			continue
		}

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
		output.WriteString("Directories:\n------------\n")
		for i, dir := range directories {
			output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, dir.Name()))
		}
	}
	
	if len(files) > 0 {
		output.WriteString("\nFiles:\n------\n")
		for i, file := range files {
			output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, file.Name()))
		}
	}

	if len(special) > 0 {
		output.WriteString("\nSpecial Files:\n--------------\n")
		for i, sp := range special {
			output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, sp.Name()))
		}
	}

	return output.String(), nil
}

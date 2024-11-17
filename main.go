package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
	"syscall"
)

func main() {

	hidden := flag.Bool("a", false, "Include hidden Files - Default is False")
	longListing := flag.Bool("l", false, "Long listing format - Default is False")
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

	output, err := goLs(inputDir, *hidden, *longListing)
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

func goLs(dir string, hidden bool, longListing bool) (string, error) {
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
			if longListing {
				info, _ := dir.Info()
				stat := info.Sys().(*syscall.Stat_t)
				owner, _ := user.LookupId(fmt.Sprint(stat.Uid))
				output.WriteString(fmt.Sprintf("%3d - %v (Owner: %s, Permissions: %o)\n", i+1, dir.Name(), owner.Username, info.Mode().Perm()))
			} else {
				output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, dir.Name()))
			}
		}
	}
	
	if len(files) > 0 {
		output.WriteString("\nFiles:\n------\n")
		for i, file := range files {
			if longListing {
				info, _ := file.Info()
				stat := info.Sys().(*syscall.Stat_t)
				sizeMB := float64(info.Size()) / (1024 * 1024)
				owner, _ := user.LookupId(fmt.Sprint(stat.Uid))
				output.WriteString(fmt.Sprintf("%3d - %v (Size: %.2f MB, Owner: %s, Permissions: %o)\n", i+1, file.Name(), sizeMB, owner.Username, info.Mode().Perm()))
			} else {
				output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, file.Name()))
			}
		}
	}

	if len(special) > 0 {
		output.WriteString("\nSpecial Files:\n--------------\n")
		for i, sp := range special {
			if longListing {
				info, _ := sp.Info()
				stat := info.Sys().(*syscall.Stat_t)
				owner, _ := user.LookupId(fmt.Sprint(stat.Uid))
				output.WriteString(fmt.Sprintf("%3d - %v (Owner: %s, Permissions: %o)\n", i+1, sp.Name(), owner.Username, info.Mode().Perm()))
			} else {
				output.WriteString(fmt.Sprintf("%3d - %v\n", i+1, sp.Name()))
			}
		}
	}

	return output.String(), nil
}

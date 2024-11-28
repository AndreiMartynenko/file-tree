package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(output io.Writer, path string, printFiles bool) error {

	return printDir(output, path, printFiles, "", false)
}

func printDir(output io.Writer, path string, printFiles bool, prefix string, isLast bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	//Sort records in alphabetical order
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	//Filter out files if printFiles is swithced off
	if !printFiles {
		var dirsOnly []os.DirEntry
		for _, entry := range entries {
			if entry.IsDir() {
				dirsOnly = append(dirsOnly, entry)
			}
		}
		entries = dirsOnly
	}
	//Iterate over the records in the directory
	for i, entry := range entries {
		isLastEntry := i == len(entries)-1
		printEntry(output, entry, prefix, isLastEntry)

		if entry.IsDir() {
			//If this is a directory, process it recursively
			newPrefix := prefix
			if isLastEntry {
				newPrefix += "\t"
			} else {
				newPrefix += "│\t"
			}
			err = printDir(output, filepath.Join(path, entry.Name()), printFiles, newPrefix, isLastEntry)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// printEntry prints one element (file or folder) of the directory tree
func printEntry(output io.Writer, entry os.DirEntry, prefix string, isLast bool) {
	// Select the desired graphical symbol
	var branch string
	if isLast {
		branch += "└───"
	} else {
		branch += "├───"
	}

	name := entry.Name()
	if entry.IsDir() {
		fmt.Fprintf(output, "%s%s%s\n", prefix, branch, name)
	} else {
		info, err := entry.Info()
		if err != nil {
			return
		}
		size := "empty"
		if info.Size() > 0 {
			size = fmt.Sprintf("%db", info.Size())
		}
		fmt.Fprintf(output, "%s%s%s (%s)\n", prefix, branch, name, size)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path> [-f]")
		return
	}

	path := os.Args[1]
	printFiles := len(os.Args) > 2 && os.Args[2] == "-f"

	err := dirTree(os.Stdout, path, printFiles)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

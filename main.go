package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func displayTree(path string, level int) {
	// your code here
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error in reading directory", err)
		return
	}

	// Loop through all files in the directory
	for _, file := range files {
		// Print the file name
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		fmt.Println(file.Name())

		// If the file is a directory, call the function recursively
		if file.IsDir() {
			newPath := filepath.Join(path, file.Name())
			displayTree(newPath, level+1)
			// displayTree(path+"/"+file.Name(), level+1)
		}
	}

}

func main() {
	// Getting the directory from the command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}
	rootDir := os.Args[1]

	// Call the function to display the tree
	fmt.Println("Tree of files for directory", rootDir)
	displayTree(rootDir, 0)
}

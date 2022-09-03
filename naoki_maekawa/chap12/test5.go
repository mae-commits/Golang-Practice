package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// A recursive function that takes the path to scan.
// We will return any errors we encounter.
func scanDirectory(path string) error {
	// Print the current directory.
	fmt.Println(path)
	// Get a slice with the directory's contents.
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// fmt.Printf("Returning error from scanDirectory(\"%s\").", path)
		panic(err)
	}

	for _, file := range files {
		// Join the directory path and filename with a slash.
		filePath := filepath.Join(path, file.Name())
		// If this is a subdirectory...
		if file.IsDir() {
			// recurcsively call scanDirectory, this time with the subDirctory's path.
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\").", path)
				return err
			}
		} else {
			// If this is a regular file, just print its path.
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	// Kick the process off by calling scanDirectory on the top directory.
	err := scanDirectory("my_directory")
	if err != nil {
		log.Fatal(err)
	}
}

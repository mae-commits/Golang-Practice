package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
		Get a slice full of values representing the contents of "my_directory".
	*/
	files, err := ioutil.ReadDir("my_directory")
	if err != nil {
		log.Fatal(err)
	}
	// For each file in the slice ...
	for _, file := range files {
		// If this file is a directory...
		if file.IsDir() {
			// Print "Directory" and the filename.
			fmt.Println("Directory:", file.Name())
		} else {
			// Otherwise, print "File" and the filename.
			fmt.Println("File:", file.Name())
		}
	}
}

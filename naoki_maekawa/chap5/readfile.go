package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// the next line opens the text data file "data.txt" in the same directory.
	file, err := os.Open("data.txt")
	// the loop below deals with the input error.
	if err != nil {
		log.Fatal(err)
	}
	// the nextline scans from file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// the nextline closes the file.
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// the nextline detecets the scan error itself.
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

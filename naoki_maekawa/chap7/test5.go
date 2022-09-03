package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// The function "GetStrings" works as inputting the string type data from the array.
// Return values are the array and the error.
func GetStrings(fileName string) ([]string, error) {
	// The next line defines the void array "lines".
	var lines []string
	// Opens the file.
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// Scanning the file.
	scanner := bufio.NewScanner(file)
	// Scanning files by every line.
	for scanner.Scan() {
		// "line" is the text type value, inputted from each line.
		line := scanner.Text()
		lines = append(lines, line)
	}
	// err closes the file.
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	// If there are no problem, return the array "lines" and no error signature "nil."
	return lines, nil
}

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	// The next line makes the map object (key(string type), value(int type))
	counts := make(map[string]int)
	// The next line extracts indices and values from the array "lines."
	// _ is a blank identifier and reveals indices of the array "lines."
	// line is values of the array "lines."
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}
}

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
	// Get values from the function "GetStrings."
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	// If there are no problem, return the array the code input from the file.
	fmt.Println(lines)
	// The array "names" includes the elements whose types are string.
	var names []string
	// The array "counts" includes the elements whose types are int.
	var counts []int
	// Input indices and elements from the array "line".
	for _, line := range lines {
		// The boolean  "matched" distunguishes whether the value is already in the array "line."
		matched := false
		// The loop below imports indices and elements in the array "names."
		for i, name := range names {
			/*
				If the element "line"(= the candidate's name) is equal to the name "line",
				increse the value of counts of that name, "counts[i]",
				then change the boolean "matched" into true.
			*/
			if name == line {
				counts[i]++
				matched = true
			}
		}
		/*
			If the boolean "matched" is false,
			it means that the name "line" has not been in the array "names" yet.
			That is why append the candidate name "line" with the array "names",
			and add 1 in the end of teh line "counts."
		*/
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	// Extract indices and values from the array "names."
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

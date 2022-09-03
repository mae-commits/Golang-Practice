package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// the function of getting the data from the file
func GetFloats(fileName string) ([]float64, error) { // output types are float64 and error, respectively.
	var numbers []float64          // define the array "numbers" without defined size.
	file, err := os.Open(fileName) // open the file.
	if err != nil {                // detect errors.
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // input values from the flie for each line.
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil { // detect errors of scanning.
			return numbers, scanner.Err()
		}
		numbers = append(numbers, number) // append the new element of new line.
	}
	err = file.Close()
	if err != nil { // detect errors
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats("data.txt")
	if err != nil { // detect errors.
		log.Fatal(err)
	}
	var sum float64 = 0
	// get both indices and elememts in line.
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount) // calculate the average of values in the file "data.txt."
}

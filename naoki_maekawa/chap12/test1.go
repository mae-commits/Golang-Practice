package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Opens the file and returns a pointer to it,
along with any error encountered.
*/
func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

// Close the file.
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	// Instead of calling os.Open directly, call OpenFile.
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	// Instead of calling file.Close directly, call CloseFile.
	// But the file will not be closed when the for loop above has an error.
	// Add "defer" so it doesn't run until GetFloats exits.
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	// Now, this right after the call to OpenFile (and its error handling code.)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			/*
				Now, even if an error is returned here, CloseFile will still called.
			*/
			return nil, err
		}
		numbers = append(numbers, number)
	}
	/*
		CloseFile would be called if an error were returned here, too.
	*/
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	/*
		And of course, CloseFile is called if GetFloats completes normally.
	*/
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}

package main

import (
	"bufio" // buffered input/output package
	"fmt"   // format package
	"log"
	"os"
	"strconv" // convert string objects to other types of them
	"strings" //
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil { // dealing with error if the input has.
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}

/*
TrimSpace remove unwanted spaces or newline inputs.

Golang cannot compare different types of functions, so when you do that,
you should convert types for only one type.

* convert string to number is impossible,
so when you would like to do that, import the strconv package and use ParseFloat method.

*/
